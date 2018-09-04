package LongestCommonPrefix

func longestCommonPrefix(strs []string) string {
	rightMost := 0
	result := ""
	for index, s := range strs {
		if index == 0 {
			rightMost = len(s)
			result = s
			continue
		}
		for i := 0; i < len(s) && i < rightMost; i++ {
			if s[i] != result[i] {
				rightMost = i
			}
		}
		if rightMost > len(s) {
			rightMost = len(s)
		}
	}
	return result[0:rightMost]
}
