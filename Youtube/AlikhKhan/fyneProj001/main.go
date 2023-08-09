package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	myApp := app.New()
	win := myApp.NewWindow("My new App Title")

	hello := widget.NewLabel("Hello Wolrd")

	win.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi", func() {
			hello.SetText("Welcome (^>.<^)")
		}),
	))

	win.ShowAndRun()
}
