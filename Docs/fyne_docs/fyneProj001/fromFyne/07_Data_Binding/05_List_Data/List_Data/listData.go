package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("List Data")

	data := binding.BindStringList(
		&[]string{"Item 1", "Item 2", "Item 3"},
	)

	list := widget.NewListWithData(
		data,

		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},

		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		},
	)

	add := widget.NewButton("Append", func() {
		val := fmt.Sprintf("Item %d", data.Length()+1)
		data.Append(val)
	})

	w.SetContent(container.NewBorder(nil, add, nil, nil, list))
	w.Resize(fyne.NewSize(250, 250))
	w.ShowAndRun()

}
