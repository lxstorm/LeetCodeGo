package SearchInRotatedSortedArray

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			if nums[mid] >= nums[l] && target < nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if nums[mid] < nums[l] && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}
