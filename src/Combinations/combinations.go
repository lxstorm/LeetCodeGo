package Combinations

func combine(n int, k int) [][]int {
	ret := [][]int{}
	mem := make([]bool, n+1)
	combination := []int{}
	var helper func(begin, step int)
	helper = func(begin, step int) {
		if step == k {
			ret = append(ret, append([]int(nil), combination...))
		} else {
			for j := begin; j <= n; j++ {
				if mem[j] == true {
					continue
				}
				combination = append(combination, j)
				mem[j] = true
				helper(j+1, step+1)
				combination = combination[:len(combination)-1]
				mem[j] = false
			}
		}
	}
	helper(1, 0)
	return ret
}
