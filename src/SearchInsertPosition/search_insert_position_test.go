package SearchInsertPosition

import "testing"

type testCase struct {
	nums   []int
	target int
	ans    int
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{
			nums:   []int{1, 3, 5, 6},
			target: 5,
			ans:    2,
		},
		testCase{
			nums:   []int{1, 3, 5, 6},
			target: 2,
			ans:    1,
		},
		testCase{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			ans:    4,
		},
		testCase{
			nums:   []int{1, 3, 5, 6},
			target: 0,
			ans:    0,
		},
	}
	for _, curCase := range testCases {
		result := searchInsert(curCase.nums, curCase.target)
		if result != curCase.ans {
			t.Errorf("Test case is %v, result is %v",
				curCase,
				result)
		}
	}

}
