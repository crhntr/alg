package fibonacci

import "testing"

func TestMatrixFibonacci(t *testing.T) {
	for nth := 0; nth < 30; nth++ {
		nthFibMatrix := MatrixFibonacci(nth)
		nthFibRecursive := RecursiveFibonacci(nth)
		//t.Logf("%d\t%d", nthFibMatrix, nthFibRecursive)
		if nthFibMatrix != nthFibRecursive {
			t.Errorf("MatrixFibonacci(%d) = %d should equal RecursiveFibonacci(%d) = %d", nth, nthFibMatrix, nth, nthFibRecursive)
		}
	}
}

func TestMatrixFibonacciLog(t *testing.T) {
	for nth := 0; nth < 30; nth++ {
		nthFibMatrix := MatrixFibonacciLog(nth)
		nthFibRecursive := RecursiveFibonacci(nth)
		//t.Logf("%d\t%d", nthFibMatrix, nthFibRecursive)
		if nthFibMatrix != nthFibRecursive {
			t.Errorf("MatrixFibonacciLog(%d) = %d should equal RecursiveFibonacci(%d) = %d", nth, nthFibMatrix, nth, nthFibRecursive)
		}
	}
}

func TestIterativeFibonacci(t *testing.T) {
	for nth := 0; nth < 30; nth++ {
		nthFibIterative := IterativeFibonacci(nth)
		nthFibRecursive := RecursiveFibonacci(nth)
		//t.Logf("%d\t%d", nthFibMatrix, nthFibRecursive)
		if nthFibIterative != nthFibRecursive {
			t.Errorf("IterativeFibonacci(%d) = %d should equal RecursiveFibonacci(%d) = %d", nth, nthFibIterative, nth, nthFibRecursive)
		}
	}
}

func benchmarkFib(i int, fubFunc func(int) int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fubFunc(i)
	}
}

func BenchmarkIterativeFib1(b *testing.B)  { benchmarkFib(1, IterativeFibonacci, b) }
func BenchmarkIterativeFib2(b *testing.B)  { benchmarkFib(2, IterativeFibonacci, b) }
func BenchmarkIterativeFib3(b *testing.B)  { benchmarkFib(3, IterativeFibonacci, b) }
func BenchmarkIterativeFib10(b *testing.B) { benchmarkFib(10, IterativeFibonacci, b) }
func BenchmarkIterativeFib20(b *testing.B) { benchmarkFib(20, IterativeFibonacci, b) }
func BenchmarkIterativeFib40(b *testing.B) { benchmarkFib(40, IterativeFibonacci, b) }
func BenchmarkIterativeFib50(b *testing.B) { benchmarkFib(50, MatrixFibonacciLog, b) }
func BenchmarkIterativeFib60(b *testing.B) { benchmarkFib(60, MatrixFibonacciLog, b) }
func BenchmarkIterativeFib70(b *testing.B) { benchmarkFib(70, MatrixFibonacciLog, b) }
func BenchmarkIterativeFib80(b *testing.B) { benchmarkFib(80, MatrixFibonacciLog, b) }
func BenchmarkIterativeFib90(b *testing.B) { benchmarkFib(90, MatrixFibonacciLog, b) }

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

func BenchmarkMatrixLogFib1(b *testing.B)  { benchmarkFib(1, MatrixFibonacciLog, b) }
func BenchmarkMatrixLogFib2(b *testing.B)  { benchmarkFib(2, MatrixFibonacciLog, b) }
func BenchmarkMatrixLogFib3(b *testing.B)  { benchmarkFib(3, MatrixFibonacciLog, b) }
func BenchmarkMatrixLogFib10(b *testing.B) { benchmarkFib(10, MatrixFibonacciLog, b) }
func BenchmarkMatrixLogFib20(b *testing.B) { benchmarkFib(20, MatrixFibonacciLog, b) }
func BenchmarkMatrixLogFib40(b *testing.B) { benchmarkFib(40, MatrixFibonacciLog, b) }

// Since MatrixFibonacciLog does not expererience exponential growth
// we can benchmark larger values in a reasonable amount of time
func BenchmarkMatrixLogFib50(b *testing.B) { benchmarkFib(50, MatrixFibonacciLog, b) }
func BenchmarkMatrixLogFib60(b *testing.B) { benchmarkFib(60, MatrixFibonacciLog, b) }
func BenchmarkMatrixLogFib70(b *testing.B) { benchmarkFib(70, MatrixFibonacciLog, b) }
func BenchmarkMatrixLogFib80(b *testing.B) { benchmarkFib(80, MatrixFibonacciLog, b) }
func BenchmarkMatrixLogFib90(b *testing.B) { benchmarkFib(90, MatrixFibonacciLog, b) }

// note the 93 Fibonacci is greater than int64 size
