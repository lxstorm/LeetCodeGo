package RemoveElement

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	result := removeElement(nums, val)
	fmt.Println(nums[0:result])
}
