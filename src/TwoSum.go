func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		pairNum := target - v
		if pairIndex, ok := m[pairNum]; ok {
			return []int{pairIndex, k}
		}
		m[v] = k
	}
	return nil
}
