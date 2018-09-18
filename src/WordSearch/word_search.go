package WordSearch

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	r, c, w := len(board), len(board[0]), len(word)
	mem := [][]bool{}
	for i := 0; i < r; i++ {
		mem = append(mem, make([]bool, c))
	}
	var helper func(i, j, wordIndex int) bool
	helper = func(i, j, wordIndex int) bool {
		if wordIndex == w {
			return true
		}
		if i < 0 || i >= r || j < 0 || j >= c || mem[i][j] || board[i][j] != word[wordIndex] {
			return false
		}
		mem[i][j] = true
		defer func() { mem[i][j] = false }()
		res := helper(i+1, j, wordIndex+1) ||
			helper(i-1, j, wordIndex+1) ||
			helper(i, j+1, wordIndex+1) ||
			helper(i, j-1, wordIndex+1)
		mem[i][j] = false
		return res
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if helper(i, j, 0) {
				return true
			}
		}
	}

	return false
}
