package knapsack_test

import (
	"testing"

	"github.com/crhntr/alg/knapsack"
)

func TestSimple(t *testing.T) {
	res := knapsack.Simple(4, 3, 1, 1, 1, 2, 5, 1, 4, 1)
	t.Log(res)
}
