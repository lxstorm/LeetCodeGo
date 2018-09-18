package RemoveDuplicatesFromSortedArrayII

func removeDuplicates(nums []int) int {
	// let i to be the place to add num
	// invariant: 0 ~ i-1 is desired slice
	i, n := 0, 2
	for j := 0; j < len(nums); j++ {
		if i < n || nums[j] > nums[i-n] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
