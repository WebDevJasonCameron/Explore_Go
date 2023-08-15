package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("gradient")

	// gradient := canvas.NewHorizontalGradient(color.White, color.Transparent)
	gradient := canvas.NewRadialGradient(color.White, color.Transparent)

	w.SetContent(gradient)

	w.Resize(fyne.NewSize(500, 150))
	w.ShowAndRun()
}
