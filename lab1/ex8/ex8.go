package main

//Импорт нескольких пакетов
import (
	"fmt"
	"math"
)

func main() {
	var defaultFloat float32
	var defaultDouble float64 = 5.5

	fmt.Println("defaultfloat       = ", defaultFloat)
	fmt.Printf("defaultDouble (%T) = %f\n\n", defaultDouble, defaultDouble)

	fmt.Println("MAX float32        = ", math.MaxFloat32)
	fmt.Println("MIN float32        = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("MAX float64        = ", math.MaxFloat64)
	fmt.Println("MIN float64        = ", math.SmallestNonzeroFloat64, "\n")

	// Завдання.
	// 1. Створіть змінні різних типів, використовуючи короткий запис та ініціалізацію за замовчуванням. Результат вивести в консоль

	intVar := 55
	floatVar := 5.5
	stringVar := "Helloooo"
	boolVar := true
	byteVar := byte(55)

	fmt.Println("Приклад змінних, створених коротким записом:")
	fmt.Printf("intVar    = %d (Type = %T)\n", intVar, intVar)
	fmt.Printf("floatVar    = %d (Type = %T)\n", floatVar, floatVar)
	fmt.Printf("stringVar    = %d (Type = %T)\n", stringVar, stringVar)
	fmt.Printf("boolVar    = %d (Type = %T)\n", boolVar, boolVar)
	fmt.Printf("byteVar    = %d (Type = %T)\n", byteVar, byteVar)
}
