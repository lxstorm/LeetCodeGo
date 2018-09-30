package ConvertSortedArrayToBinarySearchTree

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	res := sortedArrayToBST(nums)
	fmt.Println(res)
}
