package ImplementstrStr

import (
	"math"
)

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	return rabinKarp(haystack, needle)
}

func rabinKarp(haystack string, needle string) int {
	d := 16
	q := 101
	n := len(haystack)
	m := len(needle)
	p, t := 0, 0
	h := int(math.Pow(float64(d), float64(m-1))) % q

	for i := 0; i < m; i++ {
		p = (d*p + int(needle[i])) % q
		t = (d*t + int(haystack[i])) % q
	}
	for i := 0; i < n-m+1; i++ {
		if p == t {
			j := 0
			for ; j < m; j++ {
				if haystack[i+j] != needle[j] {
					break
				}
			}
			if j == m {
				return i
			}
		}
		if i+m < n {
			t = (d*(t-int(haystack[i])*h) + int(haystack[i+m])) % q
			if t < 0 {
				t = t + q
			}
		}
	}
	return -1
}
