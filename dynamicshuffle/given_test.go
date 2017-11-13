package dynamicshuffle_test

import (
	"math/rand"
	"testing"

	. "github.com/crhntr/alg/dynamicshuffle"
)

func TestIsShuffleRecursive_000111(t *testing.T) {
	x, y, w := "000", "111", "010101"

	if !IsShuffleRecursive(w, x, y) {
		t.Fail()
	}
}

func TestIsShuffleRecursiveLongRandom(t *testing.T) {
	x, y, w := "000", "111", "010101"

	if !IsShuffleRecursive(w, x, y) {
		t.Fail()
	}
}

func TestIsShuffleRecursive_011011(t *testing.T) {
	x, y, w := "011", "011", "001111"

	if !IsShuffleRecursive(w, x, y) {
		t.Fail()
	}
}

func TestIsShuffleRecursive_101101(t *testing.T) {
	x, y, w := "101", "101", "110101"

	if !IsShuffleRecursive(w, x, y) {
		t.Fail()
	}
}

func TestIsShuffleRecursive_aaabbbabcabc(t *testing.T) {
	x, y, w := "aaa", "bbb", "abcabc"

	if IsShuffleRecursive(w, x, y) {
		t.Fail()
	}
}

var (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func benchmarkShuffle(wxy []string, f func(w, x, y string) bool, b *testing.B) {
	// b.Logf("w: %q, x: %q, y: %q\n", wxy[0], wxy[1], wxy[2])
	for n := 0; n < b.N; n++ {
		if !f(wxy[0], wxy[1], wxy[2]) {
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
func BenchmarkIsShuffleRecursive1200(b *testing.B) {
	benchmarkShuffle(generateShuffle(600, 600, alphabet), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursive12000(b *testing.B) {
	benchmarkShuffle(generateShuffle(6000, 6000, alphabet), IsShuffleRecursive, b)
}
func BenchmarkIsShuffleRecursive120000(b *testing.B) {
	benchmarkShuffle(generateShuffle(60000, 60000, alphabet), IsShuffleRecursive, b)
}

func generateShuffle(yLen, xLen int, Σ string) []string {
	x, y, w := "", "", ""
	for i := 0; i < yLen; i++ {
		index := rand.Intn(len(Σ))
		y += Σ[index : index+1]
	}

	for i := 0; i < xLen; i++ {
		index := rand.Intn(len(Σ))
		x += Σ[index : index+1]
	}

	for xi, yi := 0, 0; xi+yi < (yLen + xLen); {
		if rand.Intn(2) == 1 {
			w += x[xi : xi+1]
			xi++
		} else {
			w += y[yi : yi+1]
			yi++
		}

		if xi == xLen {
			w += y[yi:]
			break
		}
		if yi == yLen {
			w += x[xi:]
			break
		}
	}

	return []string{w, x, y}
}
