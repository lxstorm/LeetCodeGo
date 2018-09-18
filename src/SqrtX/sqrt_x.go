package SqrtX

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 || x == 2 {
		return 1
	}
	l, r := 1, x
	// invariant sqrt-x is [l,r]
	for l < r {
		mid := (l + r + 1) / 2
		curPow := mid * mid
		if curPow == x {
			return mid
		}
		if curPow < x {
			l = mid
		} else {
			r = mid - 1
		}
	}

	return l
}

// anthor binary search
// sqrt-x is in [l-1, r]
// mid = (l + r) / 2
// if mid >= x/ mid    l = mid + 1
