package SubsetsII

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	ret := [][]int{}
	combine := []int{}
	n := len(nums)
	var helper func(begin, step int)
	helper = func(begin, step int) {
		tmp := append([]int(nil), combine...)
		ret = append(ret, tmp)
		if step == n {
			return
		}
		for i := begin; i < n; i++ {
			combine = append(combine, nums[i])
			helper(i+1, i+1)
			combine = combine[:len(combine)-1]
			for ; i+1 < n && nums[i+1] == nums[i]; i++ {
			}
		}
	}
	helper(0, 0)
	return ret
}
