package PermutationSequence

import (
	"strconv"
)

func getPermutation(n int, k int) string {
	res := ""
	helpSlice := []string{}
	plusSlice := []int{}
	for i := 0; i < n; i++ {
		helpSlice = append(helpSlice, strconv.Itoa(i+1))
		if i == 0 {
			plusSlice = append(plusSlice, 1)
			continue
		}
		plusSlice = append(plusSlice, i*plusSlice[len(plusSlice)-1])
	}
	remove := func(s []string, i int) []string {
		if i < 0 {
			return s
		}
		return append(s[:i], s[i+1:]...)
	}
	for ; n > 0; n-- {
		onNum := (k - 1) / plusSlice[n-1]
		res += helpSlice[onNum]
		helpSlice = remove(helpSlice, onNum)
		k -= onNum * plusSlice[n-1]
	}
	return res
}
