package FourSum

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	// nums := []int{1, 0, -1, 0, -2, 2}
	nums := []int{0, 0, 0, 0}
	target := 0
	result := fourSum(nums, target)
	fmt.Printf("Given %v, %v, result %v",
		nums,
		target,
		result)
}
