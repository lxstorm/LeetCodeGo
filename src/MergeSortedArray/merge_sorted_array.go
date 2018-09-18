package MergeSortedArray

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	l := len(nums1)
	p1, p2, i := m-1, n-1, 0
	for ; p1 >= 0 && p2 >= 0; i++ {
		if nums1[p1] > nums2[p2] {
			nums1[l-1-i] = nums1[p1]
			p1--
		} else {
			nums1[l-1-i] = nums2[p2]
			p2--
		}
	}
	for ; p1 >= 0; p1-- {
		nums1[l-1-i] = nums1[p1]
		i++
	}
	for ; p2 >= 0; p2-- {
		nums1[l-1-i] = nums2[p2]
		i++
	}
}
