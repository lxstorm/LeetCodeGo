func lengthOfLongestSubstring(s string) int {
	start, end, maxLen := 0, 0, 0
	m := map[byte]int{}

	for ; end < len(s); end++ {
		id, ok := m[s[end]]
		if ok && id >= start {
			start = id + 1
		}
		m[s[end]] = end
		curLen := end - start + 1
		if curLen > maxLen {
			maxLen = curLen
		}
	}

	return maxLen
}

//2018.2.2 rewrite
//func lengthOfLongestSubstring(s string) int {
//slen := len(s)
//m := make(map[byte]int)
//left, right, maxLen := 0, 0, 0

//for ; right < slen; right++ {
//curChar := s[right]
//lastPos, ok := m[curChar]
//m[curChar] = right
//if ok && lastPos >= left {
//left = lastPos + 1
//continue
//}
//l := right - left + 1
//if l > maxLen {
//maxLen = l
//}
//}
//return maxLen
//}
