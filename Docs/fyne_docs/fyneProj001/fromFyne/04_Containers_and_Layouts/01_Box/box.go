package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	a := app.New()
	w := a.NewWindow("Box Layout")

	t1 := canvas.NewText("Hello", color.White)
	t2 := canvas.NewText("There", color.White)
	t3 := canvas.NewText("(right)", color.White)
	content := container.New(layout.NewHBoxLayout(), t1, t2, layout.NewSpacer(), t3)

	t4 := canvas.NewText("centered", color.White)

	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), t4, layout.NewSpacer())

	w.SetContent(container.New(layout.NewVBoxLayout(), content, centered))
	w.Resize(fyne.NewSize(500, 100))
	w.ShowAndRun()
}
