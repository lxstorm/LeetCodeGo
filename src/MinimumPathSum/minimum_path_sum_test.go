package MinimumPathSum

import "testing"

type testCase struct {
	grid [][]int
	ans  int
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{
			grid: [][]int{
				[]int{1, 3, 1},
				[]int{1, 5, 1},
				[]int{4, 2, 1},
			},
			ans: 7,
		},
	}
	for _, curCase := range testCases {
		result := minPathSum(curCase.grid)
		if result != curCase.ans {
			t.Errorf("Test case is %v, result is %v", curCase, result)
		}
	}

}
