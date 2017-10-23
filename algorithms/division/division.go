package division

func Division(x, y int) (q, r int) {
	q, r = 0, x
	for y <= r {
		r, q = r-y, q+1
	}
	return
}
