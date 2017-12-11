package lms_test

import (
	"testing"

	lms "github.com/crhntr/alg/longestmonotonesequence"
)

func TestLen4(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	if ln := lms.Len(a); ln != len(a) {
		t.Errorf("the lms of a %v should be %d instead got %d", a, len(a), ln)
	}
}

func TestExampleFromBook(t *testing.T) {
	a := []int{4, 6, 5, 9, 1}
	if ln := lms.Len(a); ln != 3 {
		t.Errorf("the lms of a %v should be %d instead got %d", a, 3, ln)
	}
}
