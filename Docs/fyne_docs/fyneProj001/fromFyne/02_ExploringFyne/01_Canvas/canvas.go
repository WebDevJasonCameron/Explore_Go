package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func setContentToText(c fyne.Canvas) {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText("Text", green)
	text.TextStyle.Bold = true
	c.SetContent(text)
}

func setContenttoCircle(c fyne.Canvas) {
	red := color.NRGBA{R: 180, G: 0, B: 0, A: 255}
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	c.SetContent(circle)
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas()

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}

	rect := canvas.NewRectangle(blue)
	myCanvas.SetContent(rect)

	// go func() {
	// 	time.Sleep(time.Second)
	// 	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	// 	rect.FillColor = green
	// 	rect.Refresh()
	// }()

	go func() {
		for i := 0; i <= 2; i++ {
			if i == 1 {
				time.Sleep(time.Second)
				setContentToText(myCanvas)
			}
			if i == 2 {
				time.Sleep(time.Second)
				setContenttoCircle(myCanvas)
			}
		}

		rect.Refresh()
	}()

	myWindow.Resize(fyne.NewSize(150, 150))
	myWindow.ShowAndRun()

}
