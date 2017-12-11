package floydwarshall_test

import (
	"testing"

	"github.com/crhntr/alg/floydwarshall"
)

func TestSolve4Vertex(t *testing.T) {
	graph := NewGraph(4)

	graph.Set(0, 1, 1)
	graph.Set(1, 2, 1)
	graph.Set(2, 3, 1)
	graph.Set(2, 0, 1)
	graph.Set(3, 1, 1)

	t.Log(graph)

	floydwarshall.Solve([][]float64(graph))

	t.Log(graph)

	if !graph.IsConnected() {
		t.Error("this graph should be conected after Floyd-Warshall Alg.")
	}
}
