package palendromes

import "testing"

func TestIsPalendromeRacecar(t *testing.T) {
	if !IsPalendrome("racecar") {
		t.Fail()
	}
}

func TestIsPalendromeAsdffdsa(t *testing.T) {
	if !IsPalendrome("asdffdsa") {
		t.Fail()
	}
}

func TestIsPalendromeQwertyterwq(t *testing.T) {
	if IsPalendrome("qwertyterwq") {
		t.Fail()
	}
}
