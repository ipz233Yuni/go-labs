package main

import (
	"fmt"
	"math"
)

const (
	a    = 2147483629
	c    = 2147483587
	m    = 2147483647
	nMin = 0
	nMax = 149
	K    = 40000
)

func generateIntegers(seed int64) []int {
	values := make([]int, K)
	x := seed
	for i := 0; i < K; i++ {
		x = (a*x + c) % m
		values[i] = int(nMin + (x % int64(nMax-nMin+1)))
	}
	return values
}

func analyze(values []int) {
	// Частота інтервалів та імовірність появи випадкових величин
	freq := make(map[int]int)
	for _, v := range values {
		freq[v]++
	}

	prob := make(map[int]float64)
	for k, v := range freq {
		prob[k] = float64(v) / float64(len(values))
	}

	// Математичне сподівання
	var mean float64
	for _, v := range values {
		mean += float64(v)
	}
	mean /= float64(len(values))

	// Дисперсія
	var variance float64
	for _, v := range values {
		diff := float64(v) - mean
		variance += diff * diff
	}
	variance /= float64(len(values))
	stdDev := math.Sqrt(variance)

	fmt.Printf("=== Аналіз цілочислової послідовності ===\n")
	fmt.Printf("Кількість елементів: %d\n", len(values))
	fmt.Printf("Математичне сподівання: %.4f\n", mean)
	fmt.Printf("Дисперсія: %.4f\n", variance)
	fmt.Printf("Середньоквадратичне відхилення: %.4f\n", stdDev)

	fmt.Println("\nПерші 10 згенерованих значень:", values[:10])

	fmt.Println("\n--- Частота та ймовірність (приклад перших 10 значень) ---")
	count := 0
	for i := nMin; i <= nMax; i++ {
		if val, ok := freq[i]; ok {
			fmt.Printf("%3d -> freq: %5d, prob: %.5f\n", i, val, prob[i])
			count++
			if count >= 10 {
				break
			}
		}
	}
}

func main() {
	seed := int64(12345)
	values := generateIntegers(seed)
	analyze(values)
}
