package functions

import "testing"

func TestMinOfThree(t *testing.T) {
	x := MinOfThree(1, 2, 3)
	res := 1.0
	if x != res {
		t.Errorf("Тест не пройдено! Результат %f, має бути %f", x, res)
	}
}

func TestAverageOfThree(t *testing.T) {
	x := AverageOfThree(1, 2, 3)
	res := 2.0
	if x != res {
		t.Errorf("Тест не пройдено! Результат %f, має бути %f", x, res)
	}
}

func TestSolveEquation(t *testing.T) {
	x := SolveEquation(4, 8)
	res := -2.0
	if x != res {
		t.Errorf("Тест не пройдено! Результат %f, має бути %f", x, res)
	}
}
