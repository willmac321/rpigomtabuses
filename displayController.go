package main

import (
	"image"
	"log"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/devices/v3/ssd1306"
	"periph.io/x/devices/v3/ssd1306/image1bit"
	"periph.io/x/host/v3"
)

type DisplayController struct {
	Dev    *ssd1306.Dev
	Drawer *font.Drawer
}

var (
	// all in pixels
	advance = 7
	descent = 13
	screenW = 128
	screenH = 32
)

func clear(dc *DisplayController) {
	if err := dc.Dev.Draw(image.Rect(-1*screenW, -2*screenH, screenW, screenH*2), image.Black, image.Point{}); err != nil {
		log.Fatal(err)
	}
}

func drawBusKeepOn(dc *DisplayController) {
	clear(dc)

	dc.Dev.SetDisplayStartLine(0)

	f := basicfont.Face7x13

	img := image1bit.NewVerticalLSB(image.Rect(0, 0, screenW*2, screenH))

	drawer := font.Drawer{
		Dst:  img,
		Src:  &image.Uniform{image1bit.On},
		Face: f,
		// init offset
		Dot: fixed.P(128, 0),
	}

	newStr := []string{"   ___________", "  |_|  || o o |", "  '--O-''---O-'~"}
	dY := 0

	for len(newStr) > 0 {
		drawer.Dot = fixed.P(0, dY)
		drawer.DrawString(newStr[0])
		dY += descent
		if len(newStr) > 1 {
			newStr = newStr[1:]
		} else {
			newStr = []string{}
		}
	}

	if err := dc.Dev.Draw(dc.Dev.Bounds(), img, image.Point{}); err != nil {
		log.Fatal(err)
	}

	if err := dc.Dev.Scroll(ssd1306.Left, ssd1306.FrameRate2, 0, screenH); err != nil {
		log.Fatal(err)
	}
}

func startLoading(dc *DisplayController) {
	drawBusKeepOn(dc)
}

func stopLoading(dc *DisplayController) {
	dc.Dev.StopScroll()
	clear(dc)
}

func splitStr(str string, dc *DisplayController) []string {
	maxW := dc.Dev.Bounds().Dx()
	rv := []string{}
	count := dc.Drawer.MeasureString(str)
	if len(str) > int(maxW/advance*(screenH*2)/descent)+2 {
		str = str[:int(maxW/advance*(screenH*2)/descent)-1] + "..."
	}

	for len(str) > 0 {
		if fixed.Int26_6.Ceil(count) > maxW {
			newEnd := int(float64(maxW / advance))
			rv = append(rv, str[0:newEnd])
			str = str[newEnd:]

			// recalc count with new string
			count = dc.Drawer.MeasureString(str)
		} else {
			rv = append(rv, str)
			str = ""
		}
	}

	return rv
}

func drawMessage(newStr []string, dc *DisplayController, wait int) {
	dc.Drawer.Dst = image1bit.NewVerticalLSB(image.Rect(0, 0, screenW, screenH*2))
	dY := 9

	for len(newStr) > 0 {
		dc.Drawer.Dot = fixed.P(0, dY)
		dc.Drawer.DrawString(newStr[0])
		dY += descent
		if len(newStr) > 1 {
			newStr = newStr[1:]
		} else {
			newStr = []string{}
		}
	}

	newY := uint8(0)
	dc.Dev.SetDisplayStartLine(newY)

	if err := dc.Dev.Draw(dc.Drawer.Dst.Bounds(), dc.Drawer.Dst, image.Point{}); err != nil {
		log.Fatal(err)
	}

	i := 0
	time.Sleep(500 * time.Millisecond)

	for i < wait {
		if dY > screenH {
			dc.Dev.SetDisplayStartLine(newY)
			newY = (newY + 1) % 64
		}

		time.Sleep(100 * time.Millisecond)
		i++
	}
}

func create() *DisplayController {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Open a handle to the first available I²C bus:
	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}

	opts := ssd1306.Opts{W: screenW, H: screenH * 2, Rotated: true, Sequential: true}

	// Open a handle to a ssd1306 connected on the I²C bus:
	dev, err := ssd1306.NewI2C(bus, &opts)
	if err != nil {
		log.Fatal(err)
	}

	f := basicfont.Face7x13

	drawer := &font.Drawer{
		Src:  &image.Uniform{image1bit.On},
		Face: f,
		// init offset
		Dot: fixed.P(0, 0),
	}
	return &DisplayController{dev, drawer}
}

func dispose(dc *DisplayController) {
	dc.Dev.Halt()
}
