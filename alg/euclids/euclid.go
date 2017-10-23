package euclid

func Euclid(a, b int) (n int) {
	m, n, r := a, b, a%b
	for r > 0 {
		m, n, r = n, r, m%n
	}
	return // n
}
