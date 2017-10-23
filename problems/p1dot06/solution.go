package p1dot06

func RecursiveFibonacci(nth int) int {
	if nth < 2 {
		return 1
	}
	return RecursiveFibonacci(nth-1) + RecursiveFibonacci(nth-2)
}

func MatrixFibonacci(nth int) int {
	m1 := [2][2]int{{1, 1}, {1, 0}}
	m2 := [2][2]int{{1, 1}, {1, 0}}
	for n := 0; n < nth; n++ {
		m1 = multiplyMatrix(m1, m2)
	}
	return m1[0][1]
}

func multiplyMatrix(m1, m2 [2][2]int) [2][2]int {
	var res [2][2]int
	res[0][0] = m1[0][0]*m2[0][0] + m1[0][1]*m2[1][0]
	res[0][1] = m1[0][0]*m2[0][1] + m1[0][1]*m2[1][1]
	res[1][0] = m1[1][0]*m2[0][0] + m1[0][1]*m2[1][0]
	res[1][1] = m1[1][0]*m2[0][1] + m1[0][1]*m2[1][1]
	return res
}
