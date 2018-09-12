package MaximumSubarray

import "testing"

type testCase struct {
	input []int
	ans   int
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{
			input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			ans:   6,
		},
	}
	for _, curCase := range testCases {
		result := maxSubArray(curCase.input)
		if result != curCase.ans {
			t.Errorf("Test case is %v, result is %v",
				curCase,
				result)
		}
	}
}
