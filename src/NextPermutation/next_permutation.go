package NextPermutation

func nextPermutation(nums []int) {
	i := len(nums) - 1
	for ; i-1 >= 0 && nums[i-1] >= nums[i]; i-- {
	}

	if i != 0 {
		for j := len(nums) - 1; j >= i; {
			for ; nums[j] <= nums[i-1]; j-- {
			}
			nums[j], nums[i-1] = nums[i-1], nums[j]
		}
	}
	for k, l := i, len(nums)-1; k < l; k, l = k+1, l-1 {
		nums[k], nums[l] = nums[l], nums[k]
	}
}
