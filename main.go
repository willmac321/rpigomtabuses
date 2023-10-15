package main

func main() {

	getNearBuses()
	return

	displayController := create()

	startLoading(displayController)

	str := "0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"
	newStr := splitStr(str, displayController)

	stopLoading(displayController)

	drawMessage(newStr, displayController, 500)

	dispose(displayController)
}
