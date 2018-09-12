package RotateImage

func rotate(matrix [][]int) {
	// first row sym
	r := len(matrix)
	for i, j := 0, r-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}
	// backslash sym
	for i := 0; i < r; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
