package RemoveDuplicatesFromSortedArrayII

import "testing"

type testCase struct {
	nums []int
	ans  int
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{nums: []int{1, 1, 1, 2, 2, 3}, ans: 5},
		testCase{nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, ans: 7},
	}
	for _, curCase := range testCases {
		result := removeDuplicates(curCase.nums)
		if result != curCase.ans {
			t.Errorf("Test case is %v, result is %v", curCase, result)
		}
	}
}
