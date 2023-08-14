package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("image")

	// image := canvas.NewImageFromResource(theme.FyneLogo())
	// image := canvas.NewImageFromURI(uri)
	// image := canvas.NewImageFromImage(src)
	// image := canvas.NewImageFromReader(reader, name)
	image := canvas.NewImageFromFile("/Users/jasoncameron/00_Drive/Core/Developer/Web_Sites_Developer/01 _Jason-Codes/Initial Assets/Logo/Moon/MoonPatch_v002.png")

	image.FillMode = canvas.ImageFillOriginal

	w.SetContent(image)
	w.ShowAndRun()
}
