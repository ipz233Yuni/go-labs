package main

import (
	"fmt"
)

const (
	a    = 2147483629
	c    = 2147483587
	m    = 2147483647
	nMin = 0.0
	nMax = 149.0
	K    = 40000
)

func generateFloats(seed int64) []float64 {
	values := make([]float64, K)
	x := seed
	for i := 0; i < K; i++ {
		x = (a*x + c) % m
		values[i] = float64(x)/float64(m)*(nMax-nMin) + nMin
	}
	return values
}

func main() {
	seed := int64(12345)
	values := generateFloats(seed)

	fmt.Printf("=== Псевдовипадкова послідовність дійсних чисел ===\n")
	fmt.Println("Кількість елементів:", len(values))
	fmt.Println("Перші 10 згенерованих значень: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%.6f\n", values[i])
	}
}
