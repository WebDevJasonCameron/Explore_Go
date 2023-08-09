package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/widget"
)

func main() {

	myApp := app.New()
	win := myApp.NewWindow("My new App Title")

	win.Resize(fyne.NewSize(500, 300)) // width, height
	win.SetContent(widget.NewLabel("My first Lable"))

	win.ShowAndRun()
}
