package JumpGame

import "testing"

type testCase struct {
	input []int
	ans   bool
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{
			input: []int{2, 3, 1, 1, 4},
			ans:   true,
		},
		testCase{
			input: []int{3, 2, 1, 0, 4},
			ans:   false,
		},
		testCase{
			input: []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0},
			ans:   true,
		},
	}
	for _, curCase := range testCases {
		result := canJump(curCase.input)
		if result != curCase.ans {
			t.Errorf("Test Case is %v, result is %v",
				curCase,
				result)
		}
	}
}
