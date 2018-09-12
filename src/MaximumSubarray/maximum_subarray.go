package MaximumSubarray

import (
	"math"
)

func maxSubArray(nums []int) int {
	maxSum := math.MinInt32
	m := 0
	for i := 0; i < len(nums); i++ {
		if m < 0 {
			m = nums[i]
		} else {
			m += nums[i]
		}
		if m > maxSum {
			maxSum = m
		}
	}
	return maxSum
}

// divide and conquer solution
// func maxSubArray(nums []int) int {
// 	max := func(i, j int) int {
// 		if i > j {
// 			return i
// 		}
// 		return j
// 	}
// 	var helper func(l, r int) int
// 	helper = func(l, r int) int {
// 		if l == r {
// 			return nums[l]
// 		}
// 		mid := (l + r) / 2
// 		m1 := helper(l, mid)
// 		m2 := helper(mid+1, r)

// 		maxLeftBorderSum := math.MinInt32
// 		tmp := 0
// 		for i := mid; i >= l; i-- {
// 			tmp += nums[i]
// 			maxLeftBorderSum = max(tmp, maxLeftBorderSum)
// 		}

// 		maxRightBorderSum := math.MinInt32
// 		tmp = 0
// 		for i := mid + 1; i <= r; i++ {
// 			tmp += nums[i]
// 			maxRightBorderSum = max(tmp, maxRightBorderSum)
// 		}

// 		return max(maxLeftBorderSum+maxRightBorderSum, max(m1, m2))
// 	}

// 	return helper(0, len(nums)-1)
// }
