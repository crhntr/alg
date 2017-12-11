package lms

func Len(seq []int) int {
	if len(seq) < 1 {
		return 0
	}
	lmsAtJ := make([]int, len(seq))
	for i := range lmsAtJ {
		lmsAtJ[i] = 1
	}

	for j := 1; j < len(seq); j++ {
		max := 0

		for i := 0; i < j; i++ {
			if lmsAtJ[i] > max && seq[i] <= seq[j] {
				max = lmsAtJ[i]
			}
		}

		lmsAtJ[j] = max + 1
	}

	max := 1
	for _, ln := range lmsAtJ {
		if ln > max {
			max = ln
		}
	}
	return max
}

func Dif(seq []int, difference int) int {
	if len(seq) < 1 {
		return 0
	}
	lmsAtJ := make([]int, len(seq))
	for i := range lmsAtJ {
		lmsAtJ[i] = 1
	}

	for j := 1; j < len(seq); j++ {
		max := 0

		for i := 0; i < j; i++ {

			if lmsAtJ[i] > max {
				dif := seq[i] - seq[j]
				if dif < 0 {
					dif *= -1
				}

				if dif <= difference {
					max = lmsAtJ[i]
				}
			}
		}

		lmsAtJ[j] = max + 1
	}

	max := 1
	for _, ln := range lmsAtJ {
		if ln > max {
			max = ln
		}
	}
	return max
}
