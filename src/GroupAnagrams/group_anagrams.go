package GroupAnagrams

import "strconv"

func groupAnagrams(strs []string) [][]string {
	ret := [][]string{}
	m := map[string][]string{}
	for i := 0; i < len(strs); i++ {
		charMap := make([]int, 256)
		s := strs[i]
		for j := 0; j < len(s); j++ {
			charMap[int(s[j])]++
		}
		charIdentity := ""
		for j := 0; j < len(charMap); j++ {
			if charMap[j] != 0 {
				charIdentity += strconv.Itoa(charMap[j]) + string(j)
			}
		}
		m[charIdentity] = append(m[charIdentity], s)
	}
	for _, val := range m {
		ret = append(ret, val)
	}

	return ret

}
