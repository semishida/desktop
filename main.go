package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	a := app.New()
	w := a.NewWindow("Счёт моржи")
	w.Resize(fyne.NewSize(1280, 720))

	label := widget.NewLabel("Привет! Сколько смен ты отпахал?")
	enter := widget.NewEntry()

	label2 := widget.NewLabel("Сколько всего моржи?")
	enter2 := widget.NewEntry()

	label3 := widget.NewLabel("Сумма оклада за смену?")
	enter3 := widget.NewEntry()

	answer := widget.NewLabel("")

	btn := widget.NewButton("Посчитать ЗП!", func() {
		n1, err := strconv.ParseFloat(enter.Text, 64)
		n2, er := strconv.ParseFloat(enter2.Text, 64)
		n3, ero := strconv.ParseFloat(enter3.Text, 64)

		if err != nil || er != nil || ero != nil {
			answer.SetText("Данные неверны!")
		} else {

			sum := n1 * n3
			proc := n2 * 0.13
			itog := sum + proc

			answer.SetText(fmt.Sprintf("Итого зп:%f\n", itog))
		}
	})

	w.SetContent(container.NewVBox(
		label,
		enter,
		label2,
		enter2,
		label3,
		enter3,
		btn,
		answer,
	))
	w.ShowAndRun()

}
