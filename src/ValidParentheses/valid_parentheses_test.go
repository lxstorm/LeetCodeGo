package ValidParentheses

import "testing"

type testCase struct {
	s   string
	ans bool
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{
			s:   ")(",
			ans: false,
		},
		testCase{
			s:   "()[]{}",
			ans: true,
		},
		testCase{
			s:   "(]",
			ans: false,
		},
		testCase{
			s:   "([)]",
			ans: false,
		},
		testCase{
			s:   "{[]}",
			ans: true,
		},
	}
	for _, curCase := range testCases {
		result := isValid(curCase.s)
		if result != curCase.ans {
			t.Errorf("Test case is %v, result is %v",
				curCase,
				result)
		}
	}

}
