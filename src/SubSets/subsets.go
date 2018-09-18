package SubSets

import (
	"math"
)

func powWrapper(base, n int) int {
	return int(math.Pow(float64(base), float64(n)))
}

func subsets(nums []int) [][]int {
	n := len(nums)
	ret := [][]int{}
	for i := 0; i < powWrapper(2, n); i++ {
		tmp := []int{}
		for j := 0; j < n; j++ {
			if i>>uint(n-1-j)&0x1 != 0 {
				tmp = append(tmp, nums[j])
			}
		}
		ret = append(ret, tmp)
	}

	return ret
}

func subsetsBacktrace(nums []int) [][]int {
	n := len(nums)
	ret := [][]int{}
	subset := []int{}
	var helper func(begin, step int)
	helper = func(begin, step int) {
		ret = append(ret, append([]int(nil), subset...))
		if step == n {
			return
		}
		for i := begin; i < n; i++ {
			subset = append(subset, nums[i])
			helper(i+1, step+1)
			subset = subset[:len(subset)-1]
		}
	}
	helper(0, 0)
	return ret
}
