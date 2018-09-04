package ThreeSumCloest

import "testing"

func TestF(t *testing.T) {
	nums := []int{-1, 2, 1, -4}
	target := 1
	result := threeSumClosest(nums, target)
	t.Logf("Given %v %v, Result %v",
		nums,
		target,
		result)

}
