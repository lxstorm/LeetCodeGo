package RomanToInteger

import "testing"

type testCase struct {
	s   string
	ans int
}

func TestConvert(t *testing.T) {
	testCases := []testCase{
		testCase{
			s:   "III",
			ans: 3,
		},
		testCase{
			s:   "IV",
			ans: 4,
		},
		testCase{
			s:   "IX",
			ans: 9,
		},
		testCase{
			s:   "LVIII",
			ans: 58,
		},
		testCase{
			s:   "MCMXCIV",
			ans: 1994,
		},
	}
	for _, curCase := range testCases {
		result := romanToInt(curCase.s)
		if result != curCase.ans {
			t.Errorf("want: %v, result: %v",
				curCase.s,
				result)
		}
	}
}
