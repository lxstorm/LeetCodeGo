package CombinationSum

import (
	"fmt"
	"testing"
)

type testCase struct {
	candidates []int
	target     int
	ans        [][]int
}

func TestT(t *testing.T) {
	testCases := []testCase{
		testCase{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			ans: [][]int{
				[]int{7},
				[]int{2, 2, 3},
			},
		},
		testCase{
			candidates: []int{2, 3, 5},
			target:     8,
			ans: [][]int{
				[]int{2, 2, 2, 2},
				[]int{2, 3, 3},
				[]int{3, 5},
			},
		},
	}
	for _, curCase := range testCases {
		result := combinationSum(curCase.candidates, curCase.target)
		fmt.Println("TestCase is candidates: %v, target: %v, ans: %v, result: %v",
			curCase.candidates,
			curCase.target,
			curCase.ans,
			result)
	}
}
