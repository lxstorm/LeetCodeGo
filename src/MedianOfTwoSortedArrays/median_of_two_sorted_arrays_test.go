package MedianOfTwoSortedArrays

import "testing"

type testCase struct {
	arr1 []int
	arr2 []int
	ans  float64
}

func TestMedian(t *testing.T) {
	testCases := []testCase{
		testCase{
			arr1: []int{1, 3},
			arr2: []int{2},
			ans:  2.0,
		},
		testCase{
			arr1: []int{1, 2},
			arr2: []int{3, 4},
			ans:  2.5,
		},
		testCase{
			arr1: []int{},
			arr2: []int{3, 4, 5},
			ans:  4.0,
		},
		testCase{
			arr1: []int{2},
			arr2: []int{},
			ans:  2.0,
		},
	}
	for _, curCase := range testCases {
		result := findMedianSortedArrays(curCase.arr1, curCase.arr2)
		if result != curCase.ans {
			t.Errorf("Arr1 %v, Arr2 %v, Want %v, result is %v",
				curCase.arr1,
				curCase.arr2,
				curCase.ans,
				result)

		}

	}

}
