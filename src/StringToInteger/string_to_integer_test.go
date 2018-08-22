package StringToInteger

import "testing"

type testCase struct {
	s   string
	ans int
}

func TestAtoi(t *testing.T) {
	testCases := []testCase{
		testCase{
			s:   "42",
			ans: 42,
		},
		testCase{
			s:   "     -42",
			ans: -42,
		},
		testCase{
			s:   "4193 with words",
			ans: 4193,
		},
		testCase{
			s:   "words and 987",
			ans: 0,
		},
		testCase{
			s:   "-91283472332",
			ans: -2147483648,
		},
		testCase{
			s:   "",
			ans: 0,
		},
		testCase{
			s:   "-",
			ans: 0,
		},
	}
	for _, curCase := range testCases {
		result := myAtoi(curCase.s)
		if result != curCase.ans {
			t.Errorf("Input: %v, want: %v, result: %v",
				curCase.s,
				curCase.ans,
				result)
		}
	}

}
