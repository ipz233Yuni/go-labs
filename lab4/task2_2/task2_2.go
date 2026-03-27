package main

/*
#cgo CFLAGS: -std=c99
#include <stdbool.h>

// Оголошення C-функцій
double get_tour_price(int country, int season);
double calc_total(double days, int country, int season, bool guide, bool luxury);
*/
import "C"

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

		guide := ui.NewCheckbox("Індивідуальний гід ($50/день)")
		luxury := ui.NewCheckbox("Номер люкс (+20%)")

		result := ui.NewLabel("")
		button := ui.NewButton("Розрахувати")

		button.OnClicked(func(*ui.Button) {
			days, _ := strconv.ParseFloat(daysEntry.Text(), 64)
			c := C.int(country.Selected())
			s := C.int(season.Selected())
			g := C.bool(guide.Checked())
			l := C.bool(luxury.Checked())

			total := C.calc_total(C.double(days), c, s, g, l)
			result.SetText(fmt.Sprintf("Вартість туру: $%.2f", float64(total)))
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

		window := ui.NewWindow("Калькулятор туру (Go + C)", 500, 300, false)
		window.SetChild(grid)
		window.OnClosing(func(*ui.Window) bool { ui.Quit(); return true })
		window.Show()
	})

	if err != nil {
		panic(err)
	}
}
