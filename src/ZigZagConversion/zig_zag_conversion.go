package ZigZagConversion

import "bytes"

func nextPos(r, nRows int, dir string) int {
	var nextR int
	if dir == "up" {
		nextR = r - 1
	} else {
		nextR = r + 1
	}
	return nextR
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	store := make([]bytes.Buffer, numRows)
	// store := make([][]byte, numRows)
	res := ""
	curDir := "down"
	curR := 0
	for _, val := range s {
		store[curR].WriteByte(byte(val))
		if curR == numRows-1 {
			curDir = "up"
		}
		if curR == 0 {
			curDir = "down"
		}
		curR = nextPos(curR, numRows, curDir)
	}
	for i := range store {
		res = res + store[i].String()
	}
	return res
}
