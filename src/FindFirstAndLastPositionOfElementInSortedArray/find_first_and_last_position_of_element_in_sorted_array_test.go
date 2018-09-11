package FindFirstAndLastPositionOfElementInSortedArray

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	result := searchRange(nums, target)
	fmt.Println(result)
}
