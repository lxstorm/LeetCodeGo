package CountAndSay

import (
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := countAndSay(n - 1)
	res := ""
	memChar := byte('0')
	count := 0
	for i := 0; i < len(s); i++ {
		if i == 0 {
			memChar = byte(s[i])
			count++
		} else if memChar == byte(s[i]) {
			count++
		} else {
			res += strconv.Itoa(count) + string(memChar)
			count = 1
			memChar = byte(s[i])
		}
	}
	res += strconv.Itoa(count) + string(memChar)

	return res
}
