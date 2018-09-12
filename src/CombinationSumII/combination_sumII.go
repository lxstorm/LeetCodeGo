package CombinationSumII

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	ret := [][]int{}
	combination := []int{}
	var dfs func(begin int, target int)
	dfs = func(begin int, target int) {
		if target == 0 {
			output := append([]int(nil), combination...)
			ret = append(ret, output)
		}
		if begin >= len(candidates) || target < candidates[begin] {
			return
		}
		for i := begin; i < len(candidates); i++ {
			combination = append(combination, candidates[i])
			dfs(i+1, target-candidates[i])
			combination = combination[0 : len(combination)-1]
			for i+1 < len(candidates) && candidates[i+1] == candidates[i] {
				i++
			}
		}
	}
	sort.Slice(candidates, func(i int, j int) bool { return candidates[i] < candidates[j] })
	dfs(0, target)

	return ret

}
