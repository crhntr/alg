package dynamicshuffle

import (
	"fmt"
	"io"
)

func IsShuffleRecursive(w, u, v string) bool {
	// fmt.Printf("w: %q, u: %q, v: %q\n", w, u, v)
	if len(u) == 0 {
		return v == w
	}
	if len(v) == 0 {
		return u == w
	}
	u1, v1, w1 := u[:len(u)-1], v[:len(v)-1], w[:len(w)-1]
	return (w[len(w)-1] == v[len(v)-1] && IsShuffleRecursive(w1, u, v1)) ||
		(w[len(w)-1] == u[len(u)-1] && IsShuffleRecursive(w1, u1, v))
}

func IsShuffleRecursiveVerbose(w, u, v string, out io.Writer) bool {
	fmt.Fprintf(out, "w: %q, u: %q, v: %q\n", w, u, v)
	if len(u) == 0 {
		return v == w
	}
	if len(v) == 0 {
		return u == w
	}
	u1, v1, w1 := u[:len(u)-1], v[:len(v)-1], w[:len(w)-1]
	return (w[len(w)-1] == v[len(v)-1] && IsShuffleRecursiveVerbose(w1, u, v1, out)) ||
		(w[len(w)-1] == u[len(u)-1] && IsShuffleRecursiveVerbose(w1, u1, v, out))
}
