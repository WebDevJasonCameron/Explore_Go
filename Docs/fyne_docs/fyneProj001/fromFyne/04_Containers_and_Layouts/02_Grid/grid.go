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
	w := a.NewWindow("Grid Layout")

	t1 := canvas.NewText("1", color.White)
	t2 := canvas.NewText("2", color.White)
	t3 := canvas.NewText("3", color.White)
	grid := container.New(layout.NewGridLayout(2), t1, t2, t3)

	w.SetContent(grid)
	w.Resize(fyne.NewSize(150, 150))
	w.ShowAndRun()

}
