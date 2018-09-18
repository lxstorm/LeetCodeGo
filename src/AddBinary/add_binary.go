package AddBinary

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func addBinary(a string, b string) string {
	carry := 0
	s := ""

	for i := 0; i < len(a) || i < len(b) || carry != 0; i++ {
		aIndex, bIndex := len(a)-1-i, len(b)-1-i
		aNum, bNum := 0, 0
		if aIndex >= 0 {
			aNum = int(a[aIndex] - '0')
		}
		if bIndex >= 0 {
			bNum = int(b[bIndex] - '0')
		}
		curSum := aNum + bNum + carry
		curDigit := curSum % 2
		carry = curSum / 2
		s = string(curDigit+int('0')) + s
	}

	return s
}
