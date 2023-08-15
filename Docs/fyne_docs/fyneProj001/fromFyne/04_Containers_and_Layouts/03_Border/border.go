package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Border Layout")

	top := canvas.NewText("top bar", color.White)
	left := canvas.NewText("left", color.White)
	middle := canvas.NewText("content", color.White)
	content := container.NewBorder(top, nil, left, nil, middle)

	w.SetContent(content)
	w.Resize(fyne.NewSize(200, 200))
	w.ShowAndRun()
}
