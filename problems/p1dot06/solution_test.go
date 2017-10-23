package p1dot06

import "testing"

func TestFibonacci(t *testing.T) {
	for nth := 0; nth < 30; nth++ {
		nthFibMatrix := MatrixFibonacci(nth)
		nthFibRecursive := RecursiveFibonacci(nth)
		//t.Logf("%d\t%d", nthFibMatrix, nthFibRecursive)
		if nthFibMatrix != nthFibRecursive {
			t.Errorf("Fibonacci(nth) should equal recursiveFib(nth); however, %d != %d", nthFibMatrix, nthFibRecursive)
		}
	}
}

func benchmarkFib(i int, fubFunc func(int) int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fubFunc(i)
	}
}

func BenchmarkRecursiveFib1(b *testing.B)  { benchmarkFib(1, RecursiveFibonacci, b) }
func BenchmarkRecursiveFib2(b *testing.B)  { benchmarkFib(2, RecursiveFibonacci, b) }
func BenchmarkRecursiveFib3(b *testing.B)  { benchmarkFib(3, RecursiveFibonacci, b) }
func BenchmarkRecursiveFib10(b *testing.B) { benchmarkFib(10, RecursiveFibonacci, b) }
func BenchmarkRecursiveFib20(b *testing.B) { benchmarkFib(20, RecursiveFibonacci, b) }
func BenchmarkRecursiveFib40(b *testing.B) { benchmarkFib(40, RecursiveFibonacci, b) }

func BenchmarkMatrixFib1(b *testing.B)  { benchmarkFib(1, MatrixFibonacci, b) }
func BenchmarkMatrixFib2(b *testing.B)  { benchmarkFib(2, MatrixFibonacci, b) }
func BenchmarkMatrixFib3(b *testing.B)  { benchmarkFib(3, MatrixFibonacci, b) }
func BenchmarkMatrixFib10(b *testing.B) { benchmarkFib(10, MatrixFibonacci, b) }
func BenchmarkMatrixFib20(b *testing.B) { benchmarkFib(20, MatrixFibonacci, b) }
func BenchmarkMatrixFib40(b *testing.B) { benchmarkFib(40, MatrixFibonacci, b) }

// Since MatrixFibonacci does not expererience exponential growth
// we can benchmark larger values in a reasonable amount of time
func BenchmarkMatrixFib50(b *testing.B) { benchmarkFib(50, MatrixFibonacci, b) }
func BenchmarkMatrixFib60(b *testing.B) { benchmarkFib(60, MatrixFibonacci, b) }
func BenchmarkMatrixFib70(b *testing.B) { benchmarkFib(70, MatrixFibonacci, b) }
func BenchmarkMatrixFib80(b *testing.B) { benchmarkFib(80, MatrixFibonacci, b) }
func BenchmarkMatrixFib90(b *testing.B) { benchmarkFib(90, MatrixFibonacci, b) }

// note the 93 Fibonacci is greater than int64 size
