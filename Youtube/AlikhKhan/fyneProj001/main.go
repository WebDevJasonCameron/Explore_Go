package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("My new App Title")

	//textContent := widget.NewLabel("Hello Wolrd")

	content := widget.NewButton("click me", func() {
		log.Println("Tapped")
	})

	myWindow.SetContent(content)
	// myWindow.SetContent(container.NewVBox(
	// 	textContent,
	// 	widget.NewButton("Hi", func() {
	// 		textContent.SetText("Welcome (^>.<^)")
	// 	}),
	// ))

	myWindow.ShowAndRun()
}
