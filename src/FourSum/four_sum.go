package FourSum

import "sort"

func fourSum(nums []int, target int) [][]int {
	ret := [][]int{}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			l := j + 1
			r := len(nums) - 1
			innerTarget := target - nums[i] - nums[j]
			for l < r {
				s := nums[l] + nums[r]
				if s == innerTarget {
					tmpSlice := []int{nums[i], nums[j], nums[l], nums[r]}
					ret = append(ret, tmpSlice)
					for ; l < len(nums) && nums[l] == tmpSlice[2]; l++ {
					}
					for ; r > 0 && nums[r] == tmpSlice[3]; r-- {
					}
				} else if s < innerTarget {
					l++
				} else {
					r--
				}
			}
			for ; j+1 < len(nums) && nums[j] == nums[j+1]; j++ {
			}
		}
		for ; i+1 < len(nums) && nums[i] == nums[i+1]; i++ {
		}
	}
	return ret
}
