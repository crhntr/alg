package dynamicshuffle

import (
	"testing"
)

func TestIsShuffleRecursive_101101(t *testing.T) {
	w, u, v := "110101", "101", "101"

	if !IsShuffleRecursive(w, u, v) {
		t.Fail()
	}
}

func TestIsShuffleRecursive_0(t *testing.T) {
	w, u, v := "0", "", "0"

	if !IsShuffleRecursive(w, u, v) {
		t.Fail()
	}
}

func TestIsShuffleRecursive_1(t *testing.T) {
	w, u, v := "1", "1", ""

	if !IsShuffleRecursive(w, u, v) {
		t.Fail()
	}
}

// TestIsShuffleRecursive bad case (both sides need to be fully exauhsted)
func TestIsShuffleRecursive_BadCase(t *testing.T) {
	wuv := generateShuffle(1000, 1000, "0")

	if !IsShuffleRecursive(wuv[0], wuv[1], wuv[2]) {
		t.Fail()
	}
}

func TestIsShuffleRecursive_aaaaccebbbbdde(t *testing.T) {
	w, u, v := "aabbababcdcdee", "aaaacce", "bbbbdde"

	if !IsShuffleRecursive(w, u, v) {
		t.Fail()
	}
}

func TestIsShuffleRecursive_aaaaccqbbbbdde(t *testing.T) {
	w, u, v := "aabbababcdcdee", "aaaaccq", "bbbbdde"

	if IsShuffleRecursive(w, u, v) {
		t.Fail()
	}
}

func TestIsShuffleRecursive_aaabbbabcabc(t *testing.T) {
	w, u, v := "abcabc", "aaa", "bbb"

	if IsShuffleRecursive(w, u, v) {
		t.Fail()
	}
}
