package SearchInRotatedSortedArray

import "testing"

type testCase struct {
	nums   []int
	target int
	ans    int
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			ans:    4,
		},
		testCase{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			ans:    -1,
		},
	}
	for _, curCase := range testCases {
		result := search(curCase.nums, curCase.target)
		if result != curCase.ans {
			t.Errorf("Test Case is %v, result is %v",
				curCase,
				result)
		}

	}
}
