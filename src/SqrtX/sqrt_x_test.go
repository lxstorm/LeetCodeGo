package SqrtX

import "testing"

type testCase struct {
	input int
	ans   int
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{input: 4, ans: 2},
		testCase{input: 8, ans: 2},
	}
	for _, curCase := range testCases {
		result := mySqrt(curCase.input)
		if result != curCase.ans {
			t.Errorf("Test case is %v, result is %v", curCase, result)
		}
	}
}
