package RemoveDuplicatesFromSortedArray

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	nums := []int{1, 1, 2, 2, 2, 2, 2}
	result := removeDuplicates(nums)
	fmt.Println(result)
	fmt.Println(nums[0:result])
}
