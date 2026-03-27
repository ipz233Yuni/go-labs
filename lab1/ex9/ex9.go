package main

import "fmt"

func main() {
	var first, second bool
	var third bool = true
	fourth := !third
	var fifth = true

	fmt.Println("first  = ", first)       // false — змінна не ініціалізована
	fmt.Println("second = ", second)      // false — змінна не ініціалізована
	fmt.Println("third  = ", third)       // true— явно задано
	fmt.Println("fourth = ", fourth)      // false — !true = false
	fmt.Println("fifth  = ", fifth, "\n") // true — явно задано

	// Заперечення (!) — змінює true на false
	fmt.Println("!true  = ", !true)        // false
	fmt.Println("!false = ", !false, "\n") // true

	// Логічне І (&&) — true, якщо обидва операнди true
	fmt.Println("true && true   = ", true && true)         // true
	fmt.Println("true && false  = ", true && false)        // false
	fmt.Println("false && false = ", false && false, "\n") // false

	// Логічне АБО (||) — true, якщо хоча б один операнд true
	fmt.Println("true || true   = ", true || true)         // true
	fmt.Println("true || false  = ", true || false)        // true
	fmt.Println("false || false = ", false || false, "\n") // false

	// Порівняння чисел
	fmt.Println("2 < 3  = ", 2 < 3)        // true
	fmt.Println("2 > 3  = ", 2 > 3)        // false
	fmt.Println("3 < 3  = ", 3 < 3)        // false
	fmt.Println("3 <= 3 = ", 3 <= 3)       // true
	fmt.Println("3 > 3  = ", 3 > 3)        // false
	fmt.Println("3 >= 3 = ", 3 >= 3)       // true
	fmt.Println("2 == 3 = ", 2 == 3)       // false
	fmt.Println("3 == 3 = ", 3 == 3)       // true
	fmt.Println("2 != 3 = ", 2 != 3)       // true
	fmt.Println("3 != 3 = ", 3 != 3, "\n") // false

	// Завдання.
	// 1. Пояснити результати операцій

}
