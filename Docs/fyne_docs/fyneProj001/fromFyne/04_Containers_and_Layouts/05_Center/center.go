package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	a := app.New()
	w := a.NewWindow("Center Layout")

	//img := canvas.NewImageFromResource(theme.FyneLogo())
	//img.FillMode = canvas.ImageFillOriginal

	img := canvas.NewImageFromFile("/Users/jasoncameron/00_Drive/Core/Developer/Web_Sites_Developer/01 _Jason-Codes/Initial Assets/Logo/Moon/MoonPatch_v002.png")
	img.FillMode = canvas.ImageFillOriginal

	text := canvas.NewText("Overlay", color.Black)
	content := container.New(layout.NewCenterLayout(), img, text)

	w.SetContent(content)
	w.ShowAndRun()

}
