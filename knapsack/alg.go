package knapsack

// Simple this does not work
func Simple(capacity int, weights ...int) []int {
	if len(weights) == 0 || capacity == 0 {
		return []int{}
	}

	table := make([]bool, len(weights))
	// for i := range table {
	// 	table[i] = make([]bool, capacity+1)
	// }
	table[0] = true

	// for i := range weights {
	// 	if weights[i] < 0 {
	// 		panic("weights must not be negative")
	// 	}
	// }

	for _, wi := range weights {
		for j := capacity; j > 0; j-- {
			if j >= wi && table[j-wi] {
				table[j] = true
			}
		}
	}

	results := []int{}
	for _, wi := range weights {
		if wi > len(table) && table[wi] {
			results = append(results, wi)
		}
	}
	return results
}
