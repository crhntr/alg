package fibonacci

func IterativeFibonacci(nth int) int {
	a, b := 0, 1
	for i := 0; i <= nth; i++ {
		a, b = b, a+b
	}
	return a
}

func RecursiveFibonacci(nth int) int {
	if nth < 2 {
		return 1
	}
	return RecursiveFibonacci(nth-2) + RecursiveFibonacci(nth-1)
}

func MatrixFibonacci(nth int) int {
	m1 := [2][2]int{{1, 1}, {1, 0}}
	m2 := [2][2]int{{1, 1}, {1, 0}}
	for n := 0; n < nth; n++ {
		m1 = matrixProduct(m1, m2)
	}
	return m1[0][1]
}

// MatrixFibonacciLog was sugested by @dylanhart
func MatrixFibonacciLog(nth int) int {
	m1 := [2][2]int{{1, 1}, {1, 0}}
	m2 := [2][2]int{{1, 1}, {1, 0}}

	for ; nth > 0; nth = nth >> 1 {
		if (nth & 1) > 0 {
			m1 = matrixProduct(m1, m2)
		}
		m2 = matrixProduct(m2, m2)
	}

	return m1[1][0]
}

func matrixProduct(m1, m2 [2][2]int) [2][2]int {
	return [2][2]int{
		{m1[0][0]*m2[0][0] + m1[0][1]*m2[1][0], m1[0][0]*m2[0][1] + m1[0][1]*m2[1][1]},
		{m1[1][0]*m2[0][0] + m1[1][1]*m2[1][0], m1[1][0]*m2[0][1] + m1[1][1]*m2[1][1]},
	}
}
