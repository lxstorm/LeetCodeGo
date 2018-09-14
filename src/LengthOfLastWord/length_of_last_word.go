package LengthOfLastWord

func lengthOfLastWord(s string) int {
	count := 0
	lastEnd := len(s) - 1
	for ; lastEnd >= 0 && byte(s[lastEnd]) == ' '; lastEnd-- {
	}
	for i := lastEnd; i >= 0; i-- {
		if byte(s[i]) == ' ' {
			break
		}
		count++
	}
	return count
}
