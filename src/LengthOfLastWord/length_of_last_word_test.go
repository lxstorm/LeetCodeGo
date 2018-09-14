package LengthOfLastWord

import "testing"

type testCase struct {
	s   string
	ans int
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{"Hello World", 5},
		testCase{"a ", 1},
	}
	for _, curCase := range testCases {
		result := lengthOfLastWord(curCase.s)
		if result != curCase.ans {
			t.Errorf("Test case is %v, result is %v",
				curCase,
				result)
		}
	}
}
