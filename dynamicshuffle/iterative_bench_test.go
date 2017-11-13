package dynamicshuffle

import "testing"

func BenchmarkIsShuffleIterative6(b *testing.B) {
	benchmarkShuffle(generateShuffle(3, 3, alphabet), IsShuffleIterative, b)
}
func BenchmarkIsShuffleIterative12(b *testing.B) {
	benchmarkShuffle(generateShuffle(6, 6, alphabet), IsShuffleIterative, b)
}
func BenchmarkIsShuffleIterative120(b *testing.B) {
	benchmarkShuffle(generateShuffle(60, 60, alphabet), IsShuffleIterative, b)
}
func BenchmarkIsShuffleIterative240(b *testing.B) {
	benchmarkShuffle(generateShuffle(120, 120, alphabet), IsShuffleIterative, b)
}
func BenchmarkIsShuffleIterative1200(b *testing.B) {
	benchmarkShuffle(generateShuffle(600, 600, alphabet), IsShuffleIterative, b)
}
func BenchmarkIsShuffleIterative2400(b *testing.B) {
	benchmarkShuffle(generateShuffle(1200, 1200, alphabet), IsShuffleIterative, b)
}
func BenchmarkIsShuffleIterative12000(b *testing.B) {
	benchmarkShuffle(generateShuffle(6000, 6000, alphabet), IsShuffleIterative, b)
}
func BenchmarkIsShuffleIterative24000(b *testing.B) {
	benchmarkShuffle(generateShuffle(12000, 12000, alphabet), IsShuffleIterative, b)
}
func BenchmarkIsShuffleIterative120000(b *testing.B) {
	benchmarkShuffle(generateShuffle(60000, 60000, alphabet), IsShuffleIterative, b)
}
