package PermutationsII

import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	ret := [][]int{}
	m := map[int]bool{}
	combination := []int{}

	var dfs func(step int)
	dfs = func(step int) {
		if step == len(nums) {
			ret = append(ret, append([]int(nil), combination...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if m[i] {
				continue
			}
			m[i] = true
			combination = append(combination, nums[i])
			dfs(step + 1)
			m[i] = false
			combination = combination[0 : len(combination)-1]
			for ; i+1 < len(nums) && nums[i+1] == nums[i]; i++ {
			}
		}
	}
	dfs(0)

	return ret
}
