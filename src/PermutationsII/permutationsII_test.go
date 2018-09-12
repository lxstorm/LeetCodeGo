package PermutationsII

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	nums := []int{1, 1, 2}
	result := permuteUnique(nums)
	fmt.Println(result)
}
