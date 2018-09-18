package SearchInRotatedSortedArrayII

import "testing"

type testCase struct {
	nums   []int
	target int
	ans    bool
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		testCase{[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		testCase{[]int{1, 1, 3, 1}, 3, true},
	}
	for _, curCase := range testCases {
		result := search(curCase.nums, curCase.target)
		if result != curCase.ans {
			t.Errorf("Test case is %v, result is %v", curCase, result)
		}
	}

}
