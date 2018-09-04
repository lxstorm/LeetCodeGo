package ThreeSum

import "testing"

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	t.Logf("Given: %v, result: %v",
		nums,
		result)
}
