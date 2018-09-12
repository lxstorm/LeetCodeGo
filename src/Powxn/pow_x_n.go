package Powxn

import (
	"math"
)

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n < 0 {
		if n == math.MaxInt32 {
			return myPow(float64(x*x), n+1) / x
		}
		n = -n
		x = 1 / x
	}
	if n%2 == 0 {
		return myPow(x*x, n/2)
	}

	return myPow(float64(x*x), (n-1)/2) * float64(x)
}
