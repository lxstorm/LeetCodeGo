package CombinationSum

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	ret := [][]int{}
	tmp := []int{}
	var dfs func(l int, partTarget int)
	dfs = func(l int, partTarget int) {
		if partTarget == 0 {
			output := make([]int, len(tmp))
			copy(output, tmp)
			ret = append(ret, output)
			return
		}
		if l >= len(candidates) || partTarget < candidates[l] {
			return
		}
		for i := l; i < len(candidates); i++ {
			tmp = append(tmp, candidates[i])
			dfs(i, partTarget-candidates[i])
			tmp = tmp[0 : len(tmp)-1]
		}

	}
	sort.Slice(candidates, func(i, j int) bool { return candidates[i] < candidates[j] })
	dfs(0, target)

	return ret
}
