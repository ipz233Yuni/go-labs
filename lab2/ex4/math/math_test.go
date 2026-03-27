package math

import "testing"

func TestAdd(t *testing.T) {
	x := Add(1, 2, -3)
	res := 0.0
	if x != res {
		t.Errorf("Тест не пройдено! Результат %f, має бути %f", x, res)
	}
}

func TestAdd1(t *testing.T) {
	x := Add(1, 2, -3)
	res := 0.0
	if x != res {
		t.Errorf("Тест не пройдено! Результат %f, має бути %f", x, res)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2, -3)
	}
}
