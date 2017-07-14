func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for id, v := range nums {
		pair_id, ok := m[target-v]
		if !ok {
			m[v] = id
			continue
		}
		return []int{pair_id, id}
	}
	return nil
}
