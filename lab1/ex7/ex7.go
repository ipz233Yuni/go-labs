package main

import "fmt"

func main() {
	variable8 := int8(127)
	variable16 := int16(16383)

	fmt.Println("Приведення типів\n")

	fmt.Printf("variable8         = %-5d = %.16b\n", variable8, variable8)
	fmt.Printf("variable16        = %-5d = %.16b\n", variable16, variable16)
	fmt.Printf("uint16(variable8) = %-5d = %.16b\n", uint16(variable8), uint16(variable8))
	fmt.Printf("uint8(variable16) = %-5d = %.16b\n", uint8(variable16), uint8(variable16))

	// Завдання.
	// 1. Створіть 2 змінні різних типів. Виконайте арифметичні операції. Результат вивести в консоль

	var a int32 = 55
	var b float64 = 5.5

	sum := float64(a) + b
	diff := float64(a) - b
	mul := float64(a) * b
	div := float64(a) / b

	fmt.Println("\nАрифметичні операції:")
	fmt.Printf("%d + %.2f = %.2f\n", a, b, sum)
	fmt.Printf("%d - %.2f = %.2f\n", a, b, diff)
	fmt.Printf("%d * %.2f = %.2f\n", a, b, mul)
	fmt.Printf("%d / %.2f = %.2f\n", a, b, div)
}
