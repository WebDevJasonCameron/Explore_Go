package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
)

func main() {

	fmt.Println("Test Fyne....")

	// Start with go mod init myapp to create a package
	// Creating a new app
	a := app.New()

	// Create a new window
	w := a.NewWindow("My new title")

	w.ShowAndRun()
}
