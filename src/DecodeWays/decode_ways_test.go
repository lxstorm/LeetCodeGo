package DecodeWays

import "testing"

type testCase struct {
	input string
	ans   int
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{"12", 2},
		testCase{"2261", 3},
		testCase{"799", 1},
	}
	for _, curCase := range testCases {
		result := numDecodings(curCase.input)
		if result != curCase.ans {
			t.Errorf("Test case is %v, result is %v", curCase, result)
		}
	}
}
