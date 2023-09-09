package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("My first app")
	label := widget.NewLabel("Fyne so fine :)")

	w.SetContent(container.NewVBox(
		label,
	))
	w.ShowAndRun()

}
