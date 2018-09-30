package RestoreIPAddresses

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	l := len(s)
	ret := []string{}
	ip := []string{}
	isValid := func(s string) bool {
		if len(s) > 1 && s[0] == '0' {
			return false
		}
		num, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		if num >= 0 && num <= 255 {
			return true
		}
		return false
	}
	var helper func(int)
	helper = func(begin int) {
		if begin == l {
			if len(ip) == 4 {
				ret = append(ret, strings.Join(ip, "."))
			}
			return
		}
		for i := begin; i < l; i++ {
			curNum := s[begin : i+1]
			// note to skip the not isValid part
			if len(ip) >= 4 || !isValid(curNum) {
				break
			}
			ip = append(ip, curNum)
			helper(i + 1)
			ip = ip[0 : len(ip)-1]
		}
	}

	helper(0)

	return ret

}
