package functions

func MinOfThree(a, b, c float64) float64 {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func AverageOfThree(a, b, c float64) float64 {
	return (a + b + c) / 3
}

func SolveEquation(a, b float64) float64 {
	if a == 0 {
		panic("a не може дорівнювати 0")
	}
	return -b / a
}
