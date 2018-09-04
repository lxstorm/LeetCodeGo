package ThreeSum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	ret := [][]int{}
	for i := 0; i < len(nums); i++ {
		target := -nums[i]
		l := i + 1
		r := len(nums) - 1
		for l < r {
			s := nums[l] + nums[r]
			if s == target {
				triplet := []int{nums[i], nums[l], nums[r]}
				ret = append(ret, triplet)
				for ; l < r && nums[l] == triplet[1]; l++ {
				}
				for ; l < r && nums[r] == triplet[2]; r-- {
				}
			} else if s > target {
				r--
			} else {
				l++
			}
		}
		for ; i+1 < len(nums)-1 && nums[i+1] == nums[i]; i++ {
		}
	}
	return ret
}
