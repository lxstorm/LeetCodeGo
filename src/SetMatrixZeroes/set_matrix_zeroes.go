package SetMatrixZeroes

func clearRow(i int, matrix [][]int, columnSize int) {
	for index := 0; index < columnSize; index++ {
		matrix[i][index] = 0
	}
}
func clearColumn(i int, matrix [][]int, rowSize int) {
	for index := 0; index < rowSize; index++ {
		matrix[index][i] = 0
	}
}

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	r, c := len(matrix), len(matrix[0])

	rowFlag, columnFlag := make([]int, r), make([]int, c)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			// attention: flag first time, clear next in order to prevent multiple 0 be ommitted
			if matrix[i][j] == 0 {
				rowFlag[i] = 1
				columnFlag[j] = 1
			}
		}
	}
	for i := 0; i < len(rowFlag); i++ {
		if rowFlag[i] == 1 {
			clearRow(i, matrix, c)
		}
	}
	for i := 0; i < len(columnFlag); i++ {
		if columnFlag[i] == 1 {
			clearColumn(i, matrix, r)
		}
	}
}

// solution use constant space
func setZeroesConstant(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	r, c := len(matrix), len(matrix[0])
	row0 := 1
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					matrix[0][j] = 0
					row0 = 0
				} else {
					matrix[i][0] = 0
					matrix[0][j] = 0
				}
			}
		}
	}
	for i := 1; i < r; i++ {
		if matrix[i][0] == 0 {
			clearRow(i, matrix, c)
		}
	}
	for i := 0; i < c; i++ {
		if matrix[0][i] == 0 {
			clearColumn(i, matrix, r)
		}
	}
	if row0 == 0 {
		clearRow(0, matrix, c)
	}
}
