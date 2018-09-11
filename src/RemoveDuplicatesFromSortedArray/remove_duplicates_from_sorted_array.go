package RemoveDuplicatesFromSortedArray

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j := 0, 0

	// can use j from 1 to reduce compare j>i
	for ; j < len(nums); j++ {
		if j > i && nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i = i + 1
		}

	}

	return i + 1

}
