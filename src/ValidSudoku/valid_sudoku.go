package ValidSudoku

func isValidSudoku(board [][]byte) bool {
	nr, nc := len(board), len(board[0])
	rowValid, columnValid, squareValid := [][]int{}, [][]int{}, [][]int{}
	for i := 0; i < nr; i++ {
		rowValid = append(rowValid, make([]int, nr))
		columnValid = append(columnValid, make([]int, nr))
		squareValid = append(squareValid, make([]int, nr))
	}
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			val := board[i][j]
			if val != '.' {
				num := val - '0' - 1
				k := i/3*3 + j/3
				if rowValid[i][num] == 1 || columnValid[j][num] == 1 || squareValid[k][num] == 1 {
					return false
				}
				rowValid[i][num] = 1
				columnValid[j][num] = 1
				squareValid[k][num] = 1
			}
		}
	}
	return true

}
