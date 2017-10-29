// Package lms includes a an implementation of
// the longest monotone sequence algorithm.
package lms

func Ints(seq []int) []int {
	l := make([][]int, len(seq)) // allocation here will effect bemchmark
	l[0] = seq[:1]

	for i := 1; i < len(seq); i++ {
		for j := 0; j < i; j++ {
			if seq[j] < seq[i] && len(l[i]) < len(l[j])+1 {
				l[i] = l[j]
			}
		}
		l[i] = append(l[i], seq[i])
	}

	maxIndex := 0
	for i := 1; i < len(l); i++ {
		if len(l[i]) > len(l[maxIndex]) {
			maxIndex = i
		}
	}
	return l[maxIndex]
}
