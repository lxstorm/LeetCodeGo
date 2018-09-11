package FindFirstAndLastPositionOfElementInSortedArray

func searchRange(nums []int, target int) []int {
	ret := []int{-1, -1}
	if len(nums) == 0 {
		return ret
	}
	l, r := 0, len(nums)-1
	// search first position
	for l < r {
		mid := l + (r-l)/2
		if target == nums[mid] {
			r = mid
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if nums[l] != target {
		return ret
	}
	ret[0] = l
	//search last position
	r = len(nums) - 1
	for l < r {
		mid := l + (r-l)/2 + 1
		if target == nums[mid] {
			l = mid
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	ret[1] = r
	return ret

}
