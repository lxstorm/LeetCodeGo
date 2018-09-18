package SearchA2DMatrix

import "testing"

type testCase struct {
	matrix [][]int
	target int
	ans    bool
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{
			matrix: [][]int{
				[]int{1, 3, 5, 7},
				[]int{10, 11, 16, 20},
				[]int{23, 30, 34, 50},
			},
			target: 3,
			ans:    true,
		},
		testCase{
			matrix: [][]int{
				[]int{1, 3, 5, 7},
				[]int{10, 11, 16, 20},
				[]int{23, 30, 34, 50},
			},
			target: 13,
			ans:    false,
		},
	}
	for _, curCase := range testCases {
		result := searchMatrix(curCase.matrix, curCase.target)
		if result != curCase.ans {
			t.Errorf("Test case is %v, result is %v", curCase, result)
		}
	}
}
