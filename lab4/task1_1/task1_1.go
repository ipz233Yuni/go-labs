package main

import (
	"fmt"
	"strconv"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func main() {
	err := ui.Main(func() {
		widthEntry := ui.NewEntry()
		heightEntry := ui.NewEntry()

		material := ui.NewCombobox()
		material.Append("Дерево")
		material.Append("Метал")
		material.Append("Металопластик")
		material.SetSelected(0)

		glassType := ui.NewCombobox()
		glassType.Append("Однокамерний")
		glassType.Append("Двокамерний")
		glassType.SetSelected(0)

		windowsill := ui.NewCheckbox("Підвіконня")

		result := ui.NewLabel("")

		button := ui.NewButton("Розрахувати")

		button.OnClicked(func(*ui.Button) {
			width, _ := strconv.ParseFloat(widthEntry.Text(), 64)
			height, _ := strconv.ParseFloat(heightEntry.Text(), 64)
			area := width * height

			mat := material.Selected()
			glass := glassType.Selected()
			hasSill := windowsill.Checked()

			pricePerCm2 := getPrice(mat, glass)
			total := area * pricePerCm2
			if hasSill {
				total += 350
			}

			result.SetText(fmt.Sprintf("%.2f грн", total))
		})

		grid := ui.NewGrid()
		grid.SetPadded(true)

		grid.Append(ui.NewLabel("Ширина, см"), 0, 0, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
		grid.Append(widthEntry, 1, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		grid.Append(ui.NewLabel("Висота, см"), 0, 1, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
		grid.Append(heightEntry, 1, 1, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		grid.Append(ui.NewLabel("Матеріал"), 0, 2, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
		grid.Append(material, 1, 2, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		grid.Append(ui.NewLabel("Склопакет"), 2, 0, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
		grid.Append(glassType, 3, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		grid.Append(windowsill, 3, 1, 1, 1, false, ui.AlignStart, false, ui.AlignStart)

		grid.Append(result, 0, 4, 2, 1, true, ui.AlignFill, false, ui.AlignFill)
		grid.Append(button, 3, 4, 1, 1, false, ui.AlignEnd, false, ui.AlignFill)

		window := ui.NewWindow("Калькулятор склопакета", 600, 300, false)
		window.SetChild(grid)
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})

	if err != nil {
		panic(err)
	}
}

func getPrice(material, glass int) float64 {
	prices := [][]float64{
		{2.5, 3.0}, // Дерево: однокамерний, двокамерний
		{0.5, 1.0}, // Метал
		{1.5, 2.0}, // Металопластик
	}
	return prices[material][glass]
}
