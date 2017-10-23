package main

func GetInfix(prefix, postfix []int) (infix []int) {
	var (
		top         = prefix[0]
		left, right []int
	)

	p1 := indexOf(prefix, postfix[len(postfix)-2])
	p2 := indexOf(postfix, prefix[1]) + 1

	leftPrefix := prefix[1:p1]
	leftPostfix := postfix[0:p2]

	if equal(leftPrefix, leftPostfix) {
		left = leftPrefix
	} else {
		left = GetInfix(leftPrefix, leftPostfix)
	}

	rightPrefix := prefix[p1:]
	rightPostfix := postfix[p2 : len(postfix)-1]

	if equal(rightPrefix, rightPostfix) {
		right = rightPostfix
	} else {
		right = GetInfix(rightPrefix, rightPostfix)
	}

	return append(append(left, top), right...)
}

func GetPrefix(infix, postfix []int) (prefix []int) {
	if len(infix) == 0 {
		return []int{}
	}

	top := postfix[len(postfix)-1]
	p := indexOf(infix, top)
	left := GetPrefix(infix[:p], postfix[:p])
	right := GetPrefix(infix[p+1:], postfix[p:len(postfix)-1])
	return append(append([]int{top}, left...), right...)
}

func GetPostfix(infix, prefix []int) (postfix []int) {
	if len(infix) == 0 {
		return []int{}
	}
	top := prefix[0]
	p := indexOf(infix, top)
	left := GetPostfix(infix[:p], prefix[1:p+1])
	right := GetPostfix(infix[p+1:], prefix[p+1:])
	return append(append(left, right...), top)
}

// helper functions

func indexOf(arr []int, val int) int {
	for i := range arr {
		if arr[i] == val {
			return i
		}
	}
	panic("not found")
}

func equal(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
