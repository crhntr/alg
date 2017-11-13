package dynamicshuffle

import (
	"fmt"
	"time"
)

func IsShuffleIterative(w, u, v string) bool {
	if len(w) != (len(u) + len(v)) {
		return false
	}
	if u == "" {
		return v == w
	}
	if v == "" {
		return u == w
	}

	table := make([][]bool, len(u)+1)
	for i := range table {
		table[i] = make([]bool, len(v)+1)
	}

	// displayTable(table, u, v)

	table[0][0] = true
	for i := 1; i <= len(u); i++ {
		table[i][0] = table[i-1][0] && (w[i-1] == u[i-1])
		// displayTable(table, u, v)
	}

	for j := 1; j <= len(v); j++ {
		table[0][j] = table[0][j-1] && (w[j-1] == v[j-1])
		// displayTable(table, u, v)
	}

	for i := 1; i <= len(u); i++ {
		for j := 1; j <= len(v); j++ {
			table[i][j] = (w[i+j-1] == u[i-1] && table[i-1][j]) ||
				(w[i+j-1] == v[j-1] && table[i][j-1])
			// displayTable(table, u, v)
		}
	}

	if Verbose {
		displayTable(table, w, u, v)
	}

	return table[len(u)][len(v)]
}

func displayTable(table [][]bool, w, u, v string) {
	fmt.Printf("w: %q\n\n   %s\n", w, v)
	for i := range table {
		if i == 0 {
			fmt.Printf("  ")
		} else {
			fmt.Printf("%c ", u[i-1])
		}

		for j := range table[i] {
			if table[i][j] {
				fmt.Printf("*")
			} else {
				fmt.Printf("-")
			}
		}

		fmt.Println()
	}
	fmt.Println()
	time.Sleep(time.Millisecond * 10)
}
