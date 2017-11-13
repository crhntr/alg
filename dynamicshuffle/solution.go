package dynamicshuffle

func IsShuffleRecursive(w, x, y string) bool {
	// fmt.Printf("w: %q, x: %q, y: %q\n", w, x, y)
	if len(x) == 0 {
		return y == w
	}
	if len(y) == 0 {
		return x == w
	}
	x1, y1, w1 := x[:len(x)-1], y[:len(y)-1], w[:len(w)-1]
	return (w[len(w)-1] == y[len(y)-1] && IsShuffleRecursive(w1, x, y1)) ||
		(w[len(w)-1] == x[len(x)-1] && IsShuffleRecursive(w1, x1, y))
}
