package main

import (
	"context"
	"log"
	"os"

	// "os/exec"
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
	done := make(chan bool)

	ticker := time.NewTicker(60 * time.Second)

	doStuff()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
			case <-done:
				log.Println("Disposing Controller")
				dispose(displayController)
				return
			case <-ticker.C:
				doStuff()
			}
		}
	}()

	wg.Add(1)
	// this goroutine could be moved into the other worker and done via a timestamp calc on that loop, but wanted to look at signalling for goroutines
	go func() {
		defer wg.Done()
		log.Println("1 Hour start")
		time.Sleep(3600 * time.Second)
		// signal done to other goroutine
		done <- true
		log.Println("1 Hour completed")
	}()

	wg.Wait()

	log.Println("Process completed, sending shutdown")

	// lets turn the rpi0 after an hour of showing bus routes and let it rest
	cmd := exec.Command("shutdown", "-h", "now")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
