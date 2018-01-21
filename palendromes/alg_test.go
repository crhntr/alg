package palendromes_test

import (
	"testing"

	"github.com/crhntr/alg/palendromes"
)

func TestIsPalendromeRacecar(t *testing.T) {
	if !palendromes.Alg("racecar") {
		t.Fail()
	}
}

func TestIsPalendromeAsdffdsa(t *testing.T) {
	if !palendromes.Alg("asdffdsa") {
		t.Fail()
	}
}

func TestIsPalendromeQwertyterwq(t *testing.T) {
	if palendromes.Alg("qwertyterwq") {
		t.Fail()
	}
}
