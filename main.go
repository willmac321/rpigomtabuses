package main

func main() {
	displayController := create()

	startLoading(displayController)
	stopLoading(displayController)

	str := "0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"
	newStr := splitStr(str, displayController)

	drawMessage(newStr, displayController, 500)

	dispose(displayController)
}
