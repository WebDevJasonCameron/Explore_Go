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
	t4 := canvas.NewText("4", color.White)
	t5 := canvas.NewText("5", color.White)
	t6 := canvas.NewText("6", color.White)
	t7 := canvas.NewText("7", color.White)
	// grid := container.New(layout.NewGridLayout(3), t1, t2, t3, t4, t5, t6, t7)
	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(50, 50)), t1, t2, t3, t4, t5, t6, t7)

	w.SetContent(grid)
	w.Resize(fyne.NewSize(150, 150))
	w.ShowAndRun()

}
