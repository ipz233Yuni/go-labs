package main

import "fmt"

func main() {
	// Ініціалізація змінних
	var a int16 = 2
	var b int16 = 3
	var c int64 = 10

	fmt.Println("a + b        = ", a+b)
	fmt.Println("a - b        = ", a-b)
	fmt.Println("a * b        = ", a*b)
	fmt.Println("int(a / b)   = ", int(a/b), "\n")

	c--
	fmt.Println("c--     = ", c)
	c++
	fmt.Println("c++     = ", c)
	c += 10
	fmt.Println("c += 10 = ", c)
	c -= 5
	fmt.Println("c -= 5  = ", c)
	c *= 3
	fmt.Println("c *= 3  = ", c)
	c /= 7
	fmt.Println("c /= 7  = ", c)

	// Завдання.
	// 1. Виконати вираз: fmt.Println("c--     = ", c--)
	c--
	fmt.Println("c--     = ", c)
}
