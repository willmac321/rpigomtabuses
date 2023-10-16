package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var displayController *DisplayController

func doStuff() {
	clear(displayController)

	startLoading(displayController)
	getNearBusStops()

	strArr := compileBusFacts()

	stopLoading(displayController)

	for _, str := range strArr {
		newStr := splitStr(str, displayController)
		drawMessage(newStr, displayController, 150)
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	var wg sync.WaitGroup

	displayController = create()

	start := time.Now()
	ticker := time.NewTicker(60 * time.Second)

	doStuff()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				dispose(displayController)
				return
			case c := <-ticker.C:
				if c.Sub(start).Seconds() > 60 {
					ticker.Stop()
					break
				}
				doStuff()
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-time.After(1 * time.Minute):
			return
		}
	}()

	wg.Wait()

	log.Println("Process completed, sending shutdown")

	// lets turn the rpi0 after an hour of showing bus routes and let it rest
	cmd := exec.Command("shutdown", "-h", "now")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
