package main

import (
	"exercise/functions"
	"fmt"
)

func main() {
	fmt.Println("Мінімум:", functions.MinOfThree(1, 2, 3))
	fmt.Println("Середнє:", functions.AverageOfThree(1, 2, 3))
	fmt.Println("Рішення рівняння 2x + 4 = 0. x=", functions.SolveEquation(4, 8))
}
