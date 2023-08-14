package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Line")

	line := canvas.NewLine(color.White)
	line.StrokeWidth = 5
	line.Position1.X = 10
	line.Position1.Y = 2
	w.SetContent(line)

	w.Resize(fyne.NewSize(150, 150))
	w.ShowAndRun()
}
