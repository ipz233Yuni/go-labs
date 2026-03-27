package main

import (
	"fmt"
	"strconv"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func main() {
	err := ui.Main(func() {
		daysEntry := ui.NewEntry()

		country := ui.NewCombobox()
		country.Append("Болгарія")
		country.Append("Німеччина")
		country.Append("Польща")
		country.SetSelected(0)

		season := ui.NewCombobox()
		season.Append("Літо")
		season.Append("Зима")
		season.SetSelected(0)

		luxury := ui.NewCheckbox("Номер люкс (+20%)")
		guide := ui.NewCheckbox("Індивідуальний гід (+$50/день)")

		result := ui.NewLabel("")
		button := ui.NewButton("Розрахувати")

		button.OnClicked(func(*ui.Button) {
			days, _ := strconv.ParseFloat(daysEntry.Text(), 64)
			c := country.Selected()
			s := season.Selected()
			hasGuide := guide.Checked()
			isLuxury := luxury.Checked()

			basePrice := getTourPrice(c, s)
			total := basePrice * days

			if hasGuide {
				total += 50 * days
			}
			if isLuxury {
				total *= 1.2
			}

			result.SetText(fmt.Sprintf("Вартість туру: $%.2f", total))
		})

		grid := ui.NewGrid()
		grid.SetPadded(true)
		grid.Append(ui.NewLabel("Країна"), 0, 0, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
		grid.Append(country, 1, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		grid.Append(ui.NewLabel("Сезон"), 0, 1, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
		grid.Append(season, 1, 1, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		grid.Append(ui.NewLabel("Кількість днів"), 0, 2, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
		grid.Append(daysEntry, 1, 2, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

		grid.Append(guide, 0, 3, 1, 1, false, ui.AlignStart, false, ui.AlignStart)
		grid.Append(luxury, 1, 3, 1, 1, false, ui.AlignStart, false, ui.AlignStart)

		grid.Append(result, 0, 4, 2, 1, true, ui.AlignFill, false, ui.AlignFill)
		grid.Append(button, 1, 5, 1, 1, false, ui.AlignEnd, false, ui.AlignFill)

		window := ui.NewWindow("Калькулятор туру", 500, 300, false)
		window.SetChild(grid)
		window.OnClosing(func(*ui.Window) bool { ui.Quit(); return true })
		window.Show()
	})

	if err != nil {
		panic(err)
	}
}

func getTourPrice(country, season int) float64 {
	prices := [][]float64{
		{100, 150}, // Болгарія (літо, зима)
		{160, 200}, // Німеччина
		{120, 180}, // Польща
	}
	return prices[country][season]
}
