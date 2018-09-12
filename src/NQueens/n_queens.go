package NQueens

func solveNQueens(n int) [][]string {
	ret := [][]string{}
	solve := [][]byte{}
	// slash /  backslash \
	var rUse, cUse, slashUse, backSlashUse []bool
	init := func() {
		rUse = make([]bool, n)
		cUse = make([]bool, n)
		slashUse = make([]bool, 2*n-1)
		backSlashUse = make([]bool, 2*n-1)
		for i := 0; i < n; i++ {
			tmp := make([]byte, n)
			for j := 0; j < n; j++ {
				tmp[j] = '.'
			}
			solve = append(solve, tmp)
		}
	}
	isValid := func(i, j int) bool {
		if rUse[i] || cUse[j] || slashUse[i+j] || backSlashUse[i-j+n-1] {
			return false
		}
		return true
	}
	setUse := func(i, j int) {
		rUse[i] = true
		cUse[j] = true
		slashUse[i+j] = true
		backSlashUse[i-j+n-1] = true
	}
	clearUse := func(i, j int) {
		rUse[i] = false
		cUse[j] = false
		slashUse[i+j] = false
		backSlashUse[i-j+n-1] = false
	}
	formatSolve := func() []string {
		tmp := []string{}
		for i := 0; i < n; i++ {
			tmp = append(tmp, string(solve[i]))
		}
		return tmp
	}
	var dfs func(row, step int)
	dfs = func(i, step int) {
		if step == n {
			ret = append(ret, formatSolve())
			return
		}
		// do not need for i from to
		for j := 0; j < n; j++ {
			if !isValid(i, j) {
				continue
			}
			setUse(i, j)
			solve[i][j] = 'Q'
			dfs(i+1, step+1)
			solve[i][j] = '.'
			clearUse(i, j)
		}
	}
	init()
	dfs(0, 0)

	return ret
}
