package dynamicshuffle

import "testing"

const worstΣ = "0"

func BenchmarkIsShuffleIterativeWorstCase6(b *testing.B) {
	benchmarkShuffle(generateShuffle(3, 3, worstΣ), IsShuffleIterative, b)
}
func BenchmarkIsShuffleIterativeWorstCase12(b *testing.B) {
	benchmarkShuffle(generateShuffle(6, 6, worstΣ), IsShuffleIterative, b)
}
func BenchmarkIsShuffleIterativeWorstCase120(b *testing.B) {
	benchmarkShuffle(generateShuffle(60, 60, worstΣ), IsShuffleIterative, b)
}
func BenchmarkIsShuffleIterativeWorstCase240(b *testing.B) {
	benchmarkShuffle(generateShuffle(120, 120, worstΣ), IsShuffleIterative, b)
}
func BenchmarkIsShuffleIterativeWorstCase1200(b *testing.B) {
	benchmarkShuffle(generateShuffle(600, 600, worstΣ), IsShuffleIterative, b)
}
func BenchmarkIsShuffleIterativeWorstCase2400(b *testing.B) {
	benchmarkShuffle(generateShuffle(1200, 1200, worstΣ), IsShuffleIterative, b)
}
func BenchmarkIsShuffleIterativeWorstCase12000(b *testing.B) {
	benchmarkShuffle(generateShuffle(6000, 6000, worstΣ), IsShuffleIterative, b)
}

func BenchmarkIsShuffleRecursiveWorstCase6(b *testing.B) {
	benchmarkShuffle(generateShuffle(3, 3, worstΣ), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursiveWorstCase12(b *testing.B) {
	benchmarkShuffle(generateShuffle(6, 6, worstΣ), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursiveWorstCase120(b *testing.B) {
	benchmarkShuffle(generateShuffle(60, 60, worstΣ), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursiveWorstCase240(b *testing.B) {
	benchmarkShuffle(generateShuffle(120, 120, worstΣ), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursiveWorstCase1200(b *testing.B) {
	benchmarkShuffle(generateShuffle(600, 600, worstΣ), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursiveWorstCase2400(b *testing.B) {
	benchmarkShuffle(generateShuffle(1200, 1200, worstΣ), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursiveWorstCase12000(b *testing.B) {
	benchmarkShuffle(generateShuffle(6000, 6000, worstΣ), IsShuffleRecursive, b)
}
