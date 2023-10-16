package main

import (
	"time"
)

var displayController *DisplayController

func doStuff() {
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
	displayController = create()

	ticker := time.NewTicker(60 * time.Second)
	done := make(chan bool)

	doStuff()

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				doStuff()
			}
		}
	}()

	time.Sleep(1 * time.Hour)
	ticker.Stop()
	done <- true
	dispose(displayController)

	// lets turn the rpi0 after an hour of showing bus routes and let it rest

}
