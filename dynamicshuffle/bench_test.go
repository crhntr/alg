package dynamicshuffle_test

import (
	"testing"

	. "github.com/crhntr/alg/dynamicshuffle"
)

func benchmarkShuffle(wuv []string, f func(w, u, v string) bool, b *testing.B) {
	// if len(wuv[0]) < 100 {
	// 	b.Logf("v: %q, u: %q, y: %q\n", wuv[0], wuv[1], wuv[2])
	// }
	// b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if !f(wuv[0], wuv[1], wuv[2]) {
			b.Fail()
		}
	}
}

func BenchmarkIsShuffleRecursive6(b *testing.B) {
	benchmarkShuffle(generateShuffle(3, 3, alphabet), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursive12(b *testing.B) {
	benchmarkShuffle(generateShuffle(6, 6, alphabet), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursive120(b *testing.B) {
	benchmarkShuffle(generateShuffle(60, 60, alphabet), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursive240(b *testing.B) {
	benchmarkShuffle(generateShuffle(120, 120, alphabet), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursive1200(b *testing.B) {
	benchmarkShuffle(generateShuffle(600, 600, alphabet), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursive2400(b *testing.B) {
	benchmarkShuffle(generateShuffle(1200, 1200, alphabet), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursive12000(b *testing.B) {
	benchmarkShuffle(generateShuffle(6000, 6000, alphabet), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursive24000(b *testing.B) {
	benchmarkShuffle(generateShuffle(12000, 12000, alphabet), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursive120000(b *testing.B) {
	benchmarkShuffle(generateShuffle(60000, 60000, alphabet), IsShuffleRecursive, b)
}
