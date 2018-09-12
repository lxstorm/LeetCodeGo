package SearchInsertPosition

// func searchInsert(nums []int, target int) int {
// 	nums = append(nums, math.MaxInt32)
// 	l, r := 0, len(nums)-1
// 	for l < r {
// 		mid := l + (r-l)/2
// 		if nums[mid] == target {
// 			return mid
// 		} else if nums[mid] > target {
// 			r = mid
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return l
//
// }

// another method can handle duplicate
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	// invariant pos in low to high + 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
