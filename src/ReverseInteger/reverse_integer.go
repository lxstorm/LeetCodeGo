package ReverseInteger

import (
	"math"
)

func reverse(x int) int {
	if x == math.MinInt32 {
		return 0
	}
	absVal := x
	isPos := true
	if x < 0 {
		absVal = -x
		isPos = false
	}
	var result int64
	for absVal/10 != 0 || absVal%10 != 0 {
		result = result*10 + int64(absVal%10)
		absVal = absVal / 10
	}
	if !isPos {
		result = -result
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return int(result)
}
