package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Label Widget")

	content := widget.NewLabel("text")

	w.SetContent(content)
	w.Resize(fyne.NewSize(150, 150))
	w.ShowAndRun()
}
