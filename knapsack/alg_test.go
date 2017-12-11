package knapsack_test

import (
	"testing"

	"github.com/crhntr/alg/knapsack"
)

func TestSimple(t *testing.T) {
	res := knapsack.Simple(3, 1, 1, 1)

	t.Log(res)

}
