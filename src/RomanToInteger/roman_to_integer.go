package RomanToInteger

func romanToInt(s string) int {
	m := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}
	l := len(s)
	res := 0
	for i := 0; i < l; i++ {
		val := m[string(s[i])]
		nextVal := 0
		if i+1 < l {
			nextVal = m[string(s[i+1])]
		}
		if nextVal > val {
			val = -val
		}
		res += val
	}
	return res
}
