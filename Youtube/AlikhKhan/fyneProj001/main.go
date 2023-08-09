package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {

	myApp := app.New()
	myWin := myApp.NewWindow("My new title")

	myWin.Resize(fyne.NewSize(200, 400))

	myWin.ShowAndRun()
}
