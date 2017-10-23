package alg

import (
	"testing"
)

func CheckPrecondition(t *testing.T, condition bool, descriptionFormat string, args ...interface{}) {
	if condition {
		t.Logf("Pre-condition Passed: "+descriptionFormat, args...)
	} else {
		t.Fatalf("Pre-condition Failed: "+descriptionFormat, args...)
	}
}

func CheckPostcondition(t *testing.T, condition bool, descriptionFormat string, args ...interface{}) {
	if condition {
		t.Logf("Pre-condition Passed: "+descriptionFormat, args...)
	} else {
		t.Fatalf("Pre-condition Failed: "+descriptionFormat, args...)
	}
}
