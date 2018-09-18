package SearchInRotatedSortedArrayII

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}
		if nums[l] == nums[mid] && nums[r] == nums[mid] {
			l++
			r--
		} else {
			if nums[mid] > target {
				if nums[mid] >= nums[l] && target < nums[l] {
					l = mid + 1
				} else {
					r = mid - 1
				}
			} else {
				if nums[mid] <= nums[r] && target > nums[r] {
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
		}
	}
	return nums[l] == target
}
