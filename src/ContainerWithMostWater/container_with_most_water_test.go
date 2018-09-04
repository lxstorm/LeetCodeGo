package ContainerWithMostWater

import "testing"

type testCase struct {
	l   []int
	ans int
}

func TestMostWater(t *testing.T) {
	testCases := []testCase{
		testCase{
			l:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			ans: 49,
		},
	}
	for _, curCase := range testCases {
		result := maxArea(curCase.l)
		if result != curCase.ans {
			t.Errorf("Input: %v, want: %v, result: %v",
				curCase.l,
				curCase.ans,
				result)
		}
	}
}
