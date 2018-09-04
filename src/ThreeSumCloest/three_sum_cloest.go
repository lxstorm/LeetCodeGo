package ThreeSumCloest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	minDistance := math.MaxInt32
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	for i := 0; i < len(nums); i++ {
		innerTarget := target - nums[i]
		l := i + 1
		r := len(nums) - 1
		for l < r {
			s := nums[l] + nums[r]
			if s == innerTarget {
				return target
			} else if s < innerTarget {
				l++
			} else {
				r--
			}
			if math.Abs(float64(s-innerTarget)) < math.Abs(float64(target-minDistance)) {
				minDistance = s + nums[i]
			}
		}
	}
	return minDistance
}
