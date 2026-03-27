package main

import "fmt"

func main() {
	// Ініціалізація змінних
	var userinit8 uint8 = 1
	var userinit16 uint16 = 2
	var userinit64 int64 = -3
	var userautoinit = -4 // Такий варіант ініціалізації також можливий

	fmt.Println("Values: ", userinit8, userinit16, userinit64, userautoinit, "\n")

	// Короткий запис оголошення змінної
	// тільки для нових змінних
	intVar := 10

	fmt.Printf("Value = %d Type = %T\n", intVar, intVar)

	// Завдання.
	// 1. Вивести типи всіх змінних
	fmt.Printf("userinit8 -> Value = %v, Type = %T\n", userinit8, userinit8)
	fmt.Printf("userinit16 -> Value = %v, Type = %T\n", userinit16, userinit16)
	fmt.Printf("userinit64 -> Value = %v, Type = %T\n", userinit64, userinit64)
	fmt.Printf("userautoinit -> Value = %v, Type = %T\n", userautoinit, userautoinit)
	fmt.Printf("intVar -> Value = %v, Type = %T\n\n", intVar, intVar)
	// 2. Присвоїти змінній intVar змінні userinit16 і userautoinit. Результат вивести в консоль.
	intVar = int(userinit16)
	fmt.Printf("intVar after assigning userinit16 = %v, Type = %T\n", intVar, intVar)

	intVar = userautoinit
	fmt.Printf("intVar after assigning autoinit = %v, Type = %T\n", intVar, intVar)
}
