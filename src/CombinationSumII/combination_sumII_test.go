package CombinationSumII

import (
	"fmt"
	"testing"
)

type testCase struct {
	candidates []int
	target     int
	ans        [][]int
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			ans: [][]int{
				[]int{1, 7},
				[]int{1, 2, 5},
				[]int{2, 6},
				[]int{1, 1, 6},
			},
		},
	}
	for _, curCase := range testCases {
		result := combinationSum2(curCase.candidates, curCase.target)
		fmt.Printf("TestCase is candidates: %v, target: %v, ans: %v, result: %v\n",
			curCase.candidates,
			curCase.target,
			curCase.ans,
			result)
	}

}
