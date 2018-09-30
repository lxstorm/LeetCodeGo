package DecodeWays

import (
	"strconv"
)

// DP solution
func numDecodings(s string) int {
	isValidAplhabet := func(numStr string) bool {
		if len(numStr) == 0 || len(numStr) > 2 {
			return false
		}
		if len(numStr) == 1 {
			return numStr[0] != '0'
		}
		return numStr[0] == '1' || (numStr[0] == '2' && numStr[1] < '7')
	}

	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		if isValidAplhabet(s[0:1]) {
			return 1
		}
		return 0
	}
	mem := make([]int, len(s))
	if !isValidAplhabet(s[0:1]) {
		return 0
	}
	mem[0] = 1
	for i := 1; i < len(s); i++ {
		if isValidAplhabet(s[i : i+1]) {
			mem[i] += mem[i-1]
		}
		if isValidAplhabet(s[i-1 : i+1]) {
			if i == 1 {
				mem[i]++
			} else {
				mem[i] += mem[i-2]
			}
		}
	}

	return mem[len(s)-1]
}

// too slow
func numDecodingsBacktrack(s string) int {
	isValidAplhabet := func(numStr string) bool {
		n, err := strconv.Atoi(numStr)
		if err != nil {
			return false
		}
		if n >= 1 && n <= 26 {
			return true
		}
		return false
	}
	l := len(s)
	count := 0
	var helper func(begin int)
	helper = func(begin int) {
		if begin == l {
			count++
			return
		}
		for i := begin; i < l; i++ {
			curPart := s[begin : i+1]
			if !isValidAplhabet(curPart) {
				break
			}
			helper(i + 1)
		}
	}
	helper(0)
	return count

}
