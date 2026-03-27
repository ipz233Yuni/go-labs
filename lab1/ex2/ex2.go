package main

import "fmt"

func main() {
	// Усі змінні ініціалізуються значенням за замовчуванням, тобто 0
	var defaultInt8 int8
	var defaultInt16 int16
	var defaultInt32 int32
	var defaultInt64 int64
	var defaultInt int

	fmt.Println("Default values (signed): ", defaultInt8, defaultInt16, defaultInt32, defaultInt64, defaultInt)

	var defaultuInt8 uint8
	var defaultuInt16 uint16
	var defaultuInt32 uint32
	var defaultuInt64 uint64
	var defaultuInt uint

	fmt.Println("Default values (unsigned): ", defaultuInt8, defaultuInt16, defaultuInt32, defaultuInt64, defaultuInt)

	// Завдання.
	//1. Створити цілочисельну змінну (результат не відображати)
	var number int = 42

	_ = number
}
