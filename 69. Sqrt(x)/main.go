package main

func mySqrt(x int) int {
	if x < 0 {
		return 0
	}

	l, r := 0, x
	for l <= r {
		mid := (l + r) >> 1
		mul := mid * mid
		if mul == x {
			return mid
		} else if mul < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return r
}