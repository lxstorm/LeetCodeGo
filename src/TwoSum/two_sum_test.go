package TwoSum

import "testing"

type testCase struct {
	nums   []int
	target int
	result []int
}

func compareSliceEqual(left, right []int) bool {
	if (left == nil) != (right == nil) {
		return false
	}
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}

func TestCases(t *testing.T) {
	testCases := []testCase{
		testCase{
			nums:   []int{1, 2, 7, 9, 100},
			target: 8,
			result: []int{0, 2},
		},
		testCase{
			nums:   []int{1, 3, 4, 5},
			target: 6,
			result: []int{0, 3},
		},
	}
	caseNum := len(testCases)
	for i := 0; i < caseNum; i++ {
		curTestCase := testCases[i]
		ans := twoSum(curTestCase.nums, curTestCase.target)
		wantedResult := curTestCase.result
		isEqual := compareSliceEqual(ans, wantedResult)
		if !isEqual {
			t.Errorf("Error: Given %v, target is %v, result is %v, wanted is %v",
				curTestCase.nums,
				curTestCase.target,
				ans,
				curTestCase.result)
		}
	}
}
