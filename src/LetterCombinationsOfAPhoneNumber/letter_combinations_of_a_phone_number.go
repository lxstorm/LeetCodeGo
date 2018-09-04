package LetterCombinationsOfAPhoneNumber

func letterCombinations(digits string) []string {
	m := []string{
		"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}
	ret := []string{}
	tmp := []byte{}
	if len(digits) == 0 {
		return ret
	}

	var helper func(step int)
	helper = func(step int) {
		if step == len(digits) {
			ret = append(ret, string(tmp))
			return
		}
		loopChars := m[digits[step]-'0']
		for _, c := range loopChars {
			tmp = append(tmp, byte(c))
			helper(step + 1)
			tmp = tmp[0 : len(tmp)-1]
		}
	}
	helper(0)
	return ret
}
