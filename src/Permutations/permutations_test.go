package Permutations

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	nums := []int{1, 2, 3}
	result := permute(nums)
	fmt.Println(result)
}
