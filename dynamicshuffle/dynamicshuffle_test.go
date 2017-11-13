package dynamicshuffle

import (
	"math/rand"
	"testing"
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

var (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func generateShuffle(uLen, vLen int, Σ string) []string {
	w, u, v := "", "", ""
	for i := 0; i < uLen; i++ {
		index := rand.Intn(len(Σ))
		u += Σ[index : index+1]
	}

	for i := 0; i < vLen; i++ {
		index := rand.Intn(len(Σ))
		v += Σ[index : index+1]
	}

	for ui, vi := 0, 0; (ui + vi) < (uLen + vLen); {
		if rand.Intn(2) == 1 {
			if ui+1 <= len(u) {
				w += u[ui : ui+1]
				ui++
			}
		} else {
			if vi+1 <= len(v) {
				w += v[vi : vi+1]
				vi++
			}

		}

		if ui >= uLen {
			w += v[vi:]
			break
		}
		if vi >= vLen {
			w += u[ui:]
			break
		}
	}

	return []string{w, u, v}
}
