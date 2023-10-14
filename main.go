package main

import (
    "image"
    // "image/gif"
    "log"
    // "os"
    "time"
    "fmt"
    "math"

   "golang.org/x/image/font"
   "golang.org/x/image/font/basicfont"
   "golang.org/x/image/math/fixed"

    "periph.io/x/conn/v3/i2c/i2creg"
    "periph.io/x/devices/v3/ssd1306"
    "periph.io/x/devices/v3/ssd1306/image1bit"
    "periph.io/x/host/v3"
)

func splitStr(str string, drawer *font.Drawer) []string {

    rv := []string{}
    rect, count := drawer.BoundString(str)
    max := fixed.Int26_6.Floor(rect.Max.X)

    for len(str) > 0 {
        if count > rect.Max.X {
            newEnd := int(math.Min(float64(max/7), float64(len(str)-1)))
            rv = append(rv, str[0:newEnd])
            str = str[newEnd:len(str)-1]
            count = drawer.MeasureString(str)
        }
        fmt.Println(rv)

    }

    return rv
}

func main() {
    // Load all the drivers:
    if _, err := host.Init(); err != nil {
        log.Fatal(err)
    }

    // Open a handle to the first available I²C bus:
    bus, err := i2creg.Open("")
    if err != nil {
        log.Fatal(err)
    }

    opts := ssd1306.Opts{W:32, H:16, Rotated:true, Sequential: true}


    // Open a handle to a ssd1306 connected on the I²C bus:
    dev, err := ssd1306.NewI2C(bus, &opts)
    if err != nil {
        log.Fatal(err)
    }
    
    // Draw on it.
    img := image1bit.NewVerticalLSB(dev.Bounds())

    f := basicfont.Face7x13

    drawer := &font.Drawer{
      Dst:  img,
      Src:  &image.Uniform{image1bit.On},
      Face: f,
      Dot: fixed.P(0, 9),
      // Dot:  fixed.P(0, img.Bounds().Dy()-1-f.Descent),
    }
    fmt.Println(splitStr("Hello from periph!test", drawer))
    drawer.DrawString("Hello from periph!\r\n")
    if err := dev.Draw(dev.Bounds(), img, image.Point{}); err != nil {
      log.Fatal(err)
    }
    drawer.DrawString("second string")
    if err := dev.Draw(dev.Bounds(), img, image.Point{}); err != nil {
      log.Fatal(err)
    }

    time.Sleep(5 * time.Second)
    dev.Halt()
}
