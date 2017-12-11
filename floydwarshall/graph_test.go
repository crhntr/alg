package floydwarshall_test

import (
	"fmt"
	"math"
)

type Graph [][]float64

func NewGraph(size int) Graph {
	graph := make([][]float64, size)
	for i := range graph {
		graph[i] = make([]float64, size)
		for j := range graph[i] {
			if i == j {
				graph[i][j] = 0
			} else {
				graph[i][j] = math.Inf(1)
			}
		}
	}
	return graph
}

func (graph Graph) String() string {
	str := "\n"
	for i := range graph {
		for j := range graph[i] {
			str += fmt.Sprintf(" %0.1f ", graph[i][j])
		}
		str += "\n"
	}
	return str
}

func (graph Graph) Set(i, j int, cost float64) {
	graph[i][j] = cost
}

func (graph Graph) Get(i, j int) float64 {
	return graph[i][j]
}

func (graph Graph) IsConnected() bool {
	for i := range graph {
		for j := range graph[i] {
			if math.IsInf(graph[i][j], 1) {
				return false
			}
		}
	}
	return true
}
