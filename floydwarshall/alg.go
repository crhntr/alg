package floydwarshall

// Solve uses the floyd warshall alg. solve for the shortest paths
// from vertex i to j in the adjacency matrix initalized as follows:
// - Initialize the graph[][] adjacency matrix and n
// - graph[i][i] should be zero for all i.
// - graph[i][j] should be "infinity" if edge (i, j) does not exist
// - otherwise, graph[i][j] is the weight of the edge (i, j)
func Solve(graph [][]float64) {
	n := len(graph)

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if graph[i][k]+graph[k][j] < graph[i][j] {
					graph[i][j] = graph[i][k] + graph[k][j]
				}
			}
		}
	}
}
