package Permutations

// if [1, 2, 3] three slot, each slot has 3 possibility, and then...
func permute(nums []int) [][]int {
	ret := [][]int{}
	m := map[int]bool{}
	combination := []int{}
	var dfs func(count int)
	dfs = func(count int) {
		if count <= 0 {
			output := append([]int(nil), combination...)
			ret = append(ret, output)
			return
		}
		for i := 0; i < len(nums); i++ {
			if m[i] == true {
				continue
			}
			combination = append(combination, nums[i])
			m[i] = true
			dfs(count - 1)
			combination = combination[0 : len(combination)-1]
			m[i] = false
		}
	}
	dfs(len(nums))

	return ret
}
