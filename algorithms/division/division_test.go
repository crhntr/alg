package division

import "testing"

func TestDivision3_6(t *testing.T) {
	if q, r := Division(6, 3); q != 2 && r != 0 {
		t.Fail()
	}
}

func TestDivision7_2(t *testing.T) {
	if q, r := Division(7, 2); q != 3 && r != 1 {
		t.Fail()
	}
}

func BenchmarkRecursiveFib72(b *testing.B) {
	benchmarkDiv(7, 2, b)
}

func BenchmarkRecursiveFib702(b *testing.B) {
	benchmarkDiv(70, 2, b)
}

func BenchmarkRecursiveFib7002(b *testing.B) {
	benchmarkDiv(700, 2, b)
}

func benchmarkDiv(x, y int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Division(x, y)
	}
}
