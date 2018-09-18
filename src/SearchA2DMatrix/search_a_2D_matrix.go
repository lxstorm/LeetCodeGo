package SearchA2DMatrix

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rLen, cLen := len(matrix), len(matrix[0])
	convertNToMatrixVal := func(i int) int {
		return matrix[i/cLen][i%cLen]
	}
	n := rLen*cLen - 1
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		val := convertNToMatrixVal(mid)
		if val == target {
			return true
		} else if val > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return convertNToMatrixVal(l) == target
}
