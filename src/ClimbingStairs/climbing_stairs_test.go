package ClimbingStairs

import "testing"

type testCase struct {
	n   int
	ans int
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{2, 2},
		testCase{3, 3},
	}
	for _, curCase := range testCases {
		result := climbStairs(curCase.n)
		if result != curCase.ans {
			t.Errorf("Test case is %v, result is %v", curCase, result)
		}
	}
}
