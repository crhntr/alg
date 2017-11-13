package dynamicshuffle_test

import (
	"testing"

	. "github.com/crhntr/alg/dynamicshuffle"
)

func TestIsShuffleRecursive_000111(t *testing.T) {
	w, u, v := "010101", "000", "111"

	if !IsShuffleRecursive(w, u, v) {
		t.Fail()
	}
}

func TestIsShuffleRecursive_011011(t *testing.T) {
	w, u, v := "001111", "011", "011"

	if !IsShuffleRecursive(w, u, v) {
		t.Fail()
	}
}
