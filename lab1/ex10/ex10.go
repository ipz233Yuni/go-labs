package main

import "fmt"

func main() {
	var chartype int8 = 'R'

	fmt.Printf("Code '%c' - %d\n", chartype, chartype)

	// Завдання.
	// 1. Вивести українську літеру 'Ї'
	var char rune = 'Ї'
	fmt.Printf("Code '%c' - %d\n", char, char)
	// 2. Пояснити призначення типу "rune"
	// Пояснення:
	// rune — це псевдонім для int32, який зручно використовувати для роботи з Unicode-символами.
	// Кожен символ в Go має числовий код (Unicode code point).
	// Наприклад: 'Ї' має код 1031.
}
