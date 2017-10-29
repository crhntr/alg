package lms_test

import (
	"math/rand"
	"testing"

	. "github.com/crhntr/alg/lms"
)

func TestInts00(t *testing.T) {
	ints := []int{4, 6, 5, 9, 1}
	if len(Ints(ints)) != 3 {
		t.Fail()
	}
}

func generateInts(n int, maxVal int) []int {
	ints := make([]int, n)
	for i := range ints {
		ints[i] = rand.Intn(maxVal)
	}
	return ints
}

func benchmarkHelper(i int, b *testing.B) {
	b.StopTimer()
	ints := generateInts(i, 100)
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		Ints(ints)
	}
}

func Benchmark1(b *testing.B)  { benchmarkHelper(1, b) }
func Benchmark2(b *testing.B)  { benchmarkHelper(2, b) }
func Benchmark3(b *testing.B)  { benchmarkHelper(3, b) }
func Benchmark10(b *testing.B) { benchmarkHelper(10, b) }
func Benchmark20(b *testing.B) { benchmarkHelper(20, b) }
func Benchmark40(b *testing.B) { benchmarkHelper(40, b) }
func Benchmark50(b *testing.B) { benchmarkHelper(50, b) }
func Benchmark60(b *testing.B) { benchmarkHelper(60, b) }
func Benchmark70(b *testing.B) { benchmarkHelper(70, b) }
func Benchmark80(b *testing.B) { benchmarkHelper(80, b) }
func Benchmark90(b *testing.B) { benchmarkHelper(90, b) }
