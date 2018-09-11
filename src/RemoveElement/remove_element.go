package RemoveElement

func removeElement(nums []int, val int) int {
	i, j := 0, 0
	for ; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i = i + 1
		}
	}

	return i

}
