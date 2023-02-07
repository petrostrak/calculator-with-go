package main

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")
	w.Resize(fyne.NewSize(400, 400))

	var value_1 int = 0
	var value_2 int = 0
	var symbol string = ""
	var total string = ""
	var entryText string = ""

	entry_1 := widget.NewEntry()
	entry_1.TextStyle = fyne.TextStyle{Bold: true}
	entry_1.SetPlaceHolder("Calculator")

	btn1 := widget.NewButton("1", func() {
		total = total + "1"
		entryText = entryText + "1"
		entry_1.SetText(fmt.Sprint(entryText))
	})

	btn2 := widget.NewButton("2", func() {
		total = total + "2"
		entryText = entryText + "2"
		entry_1.SetText(fmt.Sprint(entryText))
	})

	btn3 := widget.NewButton("3", func() {
		total = total + "3"
		entryText = entryText + "3"
		entry_1.SetText(fmt.Sprint(entryText))
	})

	btn4 := widget.NewButton("4", func() {
		total = total + "4"
		entryText = entryText + "4"
		entry_1.SetText(fmt.Sprint(entryText))
	})

	btn5 := widget.NewButton("5", func() {
		total = total + "5"
		entryText = entryText + "5"
		entry_1.SetText(fmt.Sprint(entryText))
	})

	btn6 := widget.NewButton("6", func() {
		total = total + "6"
		entryText = entryText + "6"
		entry_1.SetText(fmt.Sprint(entryText))
	})

	btn7 := widget.NewButton("7", func() {
		total = total + "7"
		entryText = entryText + "7"
		entry_1.SetText(fmt.Sprint(entryText))
	})

	btn8 := widget.NewButton("8", func() {
		total = total + "8"
		entryText = entryText + "8"
		entry_1.SetText(fmt.Sprint(entryText))
	})

	btn9 := widget.NewButton("9", func() {
		total = total + "9"
		entryText = entryText + "9"
		entry_1.SetText(fmt.Sprint(entryText))
	})

	btn0 := widget.NewButton("0", func() {
		total = total + "0"
		entryText = entryText + "0"
		entry_1.SetText(fmt.Sprint(entryText))
	})

	btnDot := widget.NewButton(".", func() {})

	result := canvas.NewText("Result", color.Black)
	result.TextSize = 30
	result.Alignment = fyne.TextAlignCenter

	btnClear := widget.NewButton("Clear", func() {
		entry_1.SetText("")
		total = ""
		result.Text = ""
		entryText = ""
	})

	btnPlus := widget.NewButton("+", func() {
		value_1, _ = strconv.Atoi(total)
		symbol = "+"
		entryText = entryText + "+"
		entry_1.SetText(fmt.Sprint(entryText))
		total = ""
	})

	btnMinus := widget.NewButton("-", func() {
		value_1, _ = strconv.Atoi(total)
		symbol = "-"
		entryText = entryText + "-"
		entry_1.SetText(fmt.Sprint(entryText))
		total = ""
	})

	btnMutli := widget.NewButton("*", func() {
		value_1, _ = strconv.Atoi(total)
		symbol = "*"
		entryText = entryText + "*"
		entry_1.SetText(fmt.Sprint(entryText))
		total = ""
	})

	btnDiv := widget.NewButton("?", func() {
		value_1, _ = strconv.Atoi(total)
		symbol = "/"
		entryText = entryText + "/"
		entry_1.SetText(fmt.Sprint(entryText))
		total = ""
	})

	btnEqual := widget.NewButton("=", func() {
		value_2, _ = strconv.Atoi(total)
		switch symbol {
		case "+":
			answer := value_1 + value_2
			result.Text = fmt.Sprint(answer)
			result.Refresh()
		case "-":
			answer := value_1 - value_2
			result.Text = fmt.Sprint(answer)
			result.Refresh()
		case "*":
			answer := value_1 * value_2
			result.Text = fmt.Sprint(answer)
			result.Refresh()
		case "/":
			answer := value_1 / value_2
			result.Text = fmt.Sprint(answer)
			result.Refresh()
		}
		total = ""
	})

	w.SetContent(
		container.NewVBox(
			entry_1, result,
			container.NewGridWithColumns(4,
				btn1,
				btn1,
				btn2,
				btn3,
				btn4,
				btn5,
				btn6,
				btn7,
				btn8,
				btn9,
				btn0,
				btnDot,
				btnEqual,
				btnClear,
				btnPlus,
				btnMinus,
				btnMutli,
				btnDiv,
			),
		),
	)
	w.ShowAndRun()
}
