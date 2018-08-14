package TwoSum

func twoSum(nums []int, target int) []int {
	// m 用来保存数字与其index的map关系
	m := make(map[int]int)
	result := []int{}
	for index, value := range nums {
		pairNum := target - value
		if pairIndex, ok := m[pairNum]; ok {
			result = []int{pairIndex, index}
			return result
		}
		m[value] = index
	}

	return result
}
