package NextPermutation

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	testCases := [][]int{
		[]int{1, 2, 3},
		[]int{1, 2, 5, 4, 3},
		[]int{1, 1},
	}
	for _, curCase := range testCases {
		nextPermutation(curCase)
		fmt.Println(curCase)
	}
}
