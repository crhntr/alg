package dynamicshuffle

import "testing"

// TestIsShuffleIterative_000111 given
func TestIsShuffleIterative_000111(t *testing.T) {
	w, u, v := "010101", "000", "111"

	if !IsShuffleIterative(w, u, v) {
		t.Fail()
	}
}

// TestIsShuffleIterative_011011 given
func TestIsShuffleIterative_011011(t *testing.T) {
	w, u, v := "001111", "011", "011"

	if !IsShuffleIterative(w, u, v) {
		t.Fail()
	}
}

// TestIsShuffleIterative_011011 given
func TestIsShuffleIterative_BadCase(t *testing.T) {
	wuv := generateShuffle(1000, 1000, "0")

	if !IsShuffleIterative(wuv[0], wuv[1], wuv[2]) {
		t.Fail()
	}
}

func TestIsShuffleIterative_01(t *testing.T) {
	w, u, v := "01", "", "01"

	if !IsShuffleIterative(w, u, v) {
		t.Fail()
	}
}

func TestIsShuffleIterative_10(t *testing.T) {
	w, u, v := "10", "10", ""

	if !IsShuffleIterative(w, u, v) {
		t.Fail()
	}
}

func TestIsShuffleIterative_100(t *testing.T) {
	w, u, v := "100", "", "10"

	if IsShuffleIterative(w, u, v) {
		t.Fail()
	}
}

func TestIsShuffleIterative_aaaaccebbbbdde(t *testing.T) {
	w, u, v := "aabbababcdcdee", "aaaacce", "bbbbdde"

	if !IsShuffleIterative(w, u, v) {
		t.Fail()
	}
}

func TestIsShuffleIterative_ab(t *testing.T) {
	w, u, v := "aa", "a", "b"

	if IsShuffleIterative(w, u, v) {
		t.Fail()
	}
}

func TestIsShuffleIterative_aaabbbabcabc(t *testing.T) {
	w, u, v := "abcabc", "aaa", "bbb"

	if IsShuffleIterative(w, u, v) {
		t.Fail()
	}
}
