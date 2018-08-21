package LongestPalindromicSubstring

// DP solution
// top-down slow

// func longestPalindrome(s string) string {
// 	m := make(map[string]bool)
// 	var dpHelper func(i, j int, s string) bool
// 	dpHelper = func(i, j int, s string) bool {
// 		genKey := func(i, j int) string {
// 			key := strconv.Itoa(i) + "," + strconv.Itoa(j)
// 			return key
// 		}
// 		key := genKey(i, j)
// 		if val, ok := m[key]; ok {
// 			return val
// 		}
// 		if j-i < 3 {
// 			m[key] = s[i] == s[j]
// 			return m[key]
// 		}
// 		// attention!! dpHelper need to put first due to lazy cal
// 		m[key] = dpHelper(i+1, j-1, s) && s[i] == s[j]
// 		return m[key]
// 	}

// 	l := len(s)
// 	if l == 0 {
// 		return ""
// 	}
// 	start, end := 0, 0

// 	for i := 0; i < l; i++ {
// 		for j := i; j < l; j++ {
// 			isPalin := dpHelper(i, j, s)
// 			if isPalin {
// 				if (j - i + 1) >= end-start+1 {
// 					start = i
// 					end = j
// 				}
// 			}
// 		}
// 	}
// 	return s[start : end+1]

// }

// use expand centor method
func getMaxPalindrome(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// right - 1 - (left + 1) + 1
	return right - left - 1
}

func longestPalindrome(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < l; i++ {
		// another solution don't calculate two length
		// if abbc, now at first b, jump i to second b, otherwise we can't
		// include bb
		len1 := getMaxPalindrome(s, i, i+1)
		len2 := getMaxPalindrome(s, i, i)
		length := len1
		if len1 < len2 {
			length = len2
		}
		if length >= end-start+1 {
			start = i - (length-1)/2
			end = i + length/2
		}
	}
	return s[start : end+1]
}
