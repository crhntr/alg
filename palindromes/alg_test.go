package palindromes

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		Value  string
		Expect bool
	}{
		{Value: "racecar", Expect: true},
		{Value: "asdffdsa", Expect: true},
		{Value: "qwertyterwq", Expect: false},
	}
	for _, test := range tests {
		t.Run(test.Value, func(t *testing.T) {
			if Alg(test.Value) != test.Expect {
				t.Fail()
			}
		})
	}
}
