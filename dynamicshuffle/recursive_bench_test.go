package dynamicshuffle

import (
	"testing"
)

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
