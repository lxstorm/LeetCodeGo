package MergeSortedArray

import (
	"fmt"
	"testing"
)

type testCase struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
}

func TestF(t *testing.T) {
	// nums1 := []int{1, 2, 3, 0, 0, 0, 0, 0, 0}
	// nums2 := []int{2, 5, 6}
	// m, n := 3, 3
	nums1 := []int{0}
	nums2 := []int{1}
	m, n := 0, 1
	merge(nums1, m, nums2, n)
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
