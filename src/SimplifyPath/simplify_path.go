package SimplifyPath

import (
	"strings"
)

func simplifyPath(path string) string {
	path = path + "/"
	pathSlice := []string{}
	realPathSlice := []string{}
	preSlashIndex := 0
	for i := 1; i < len(path); i++ {
		if path[i] == '/' && preSlashIndex < i {
			sToAdd := path[preSlashIndex+1 : i]
			if sToAdd != "" {
				pathSlice = append(pathSlice, sToAdd)
			}
			preSlashIndex = i
		}
	}
	for i := 0; i < len(pathSlice); i++ {
		curStr := pathSlice[i]
		if curStr != "." && curStr != ".." {
			realPathSlice = append(realPathSlice, curStr)
		} else if curStr == ".." {
			if len(realPathSlice) > 0 {
				realPathSlice = realPathSlice[:len(realPathSlice)-1]
			}
		}
	}

	return "/" + strings.Join(realPathSlice, "/")
}
