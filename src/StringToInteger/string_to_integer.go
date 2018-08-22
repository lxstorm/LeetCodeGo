package StringToInteger

import (
	"math"
)

func isValidateChar(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func myAtoi(str string) int {
	posFlag := 1
	index := 0
	var result int64
	// skip blank char
	for ; index < len(str) && str[index] == ' '; index++ {
	}
	if index < len(str) {
		if str[index] == '-' {
			posFlag = -1
			index++
		} else if str[index] == '+' {
			index++
		}
	}
	// begin cal
	for ; index < len(str); index++ {
		curChar := str[index]
		if !isValidateChar(curChar) {
			break
		} else {
			addon := int64(curChar) - int64('0')
			result = result*10 + int64(posFlag)*addon
			if result < math.MinInt32 {
				result = math.MinInt32
			}
			if result > math.MaxInt32 {
				result = math.MaxInt32
				break
			}
		}
	}
	return int(result)
}
