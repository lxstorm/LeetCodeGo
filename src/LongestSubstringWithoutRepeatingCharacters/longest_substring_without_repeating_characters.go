package LongestSubstringWithoutRepeatingCharacters

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	leftPos, rightPos, maxLen := 0, 0, 0
	l := len(s)
	for ; rightPos < l; rightPos++ {
		curChar := s[rightPos]
		lastPos, ok := m[curChar]
		if ok && lastPos >= leftPos {
			leftPos = lastPos + 1
		}
		// calulate the cur not repeating substr len
		// otherwise the last not repeating len
		curLen := rightPos - leftPos + 1
		if curLen > maxLen {
			maxLen = curLen
		}
		m[curChar] = rightPos
	}
	return maxLen
}
