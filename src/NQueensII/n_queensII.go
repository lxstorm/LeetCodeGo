package NQueensII

func totalNQueens(n int) int {
	count := 0
	// slash /  backslash \
	var cUse, slashUse, backSlashUse []bool
	var solve [][]byte

	init := func() {
		cUse = make([]bool, n)
		slashUse = make([]bool, 2*n-1)
		backSlashUse = make([]bool, 2*n-1)
		solve = make([][]byte, n)
		for i := 0; i < n; i++ {
			tmp := make([]byte, n)
			for j := 0; j < n; j++ {
				tmp[j] = '.'
			}
			solve[i] = tmp
		}
	}
	setUse := func(i, j int) {
		cUse[j] = true
		slashUse[i+j] = true
		backSlashUse[i-j+n-1] = true
	}
	clearUse := func(i, j int) {
		cUse[j] = false
		slashUse[i+j] = false
		backSlashUse[i-j+n-1] = false
	}
	isValid := func(i, j int) bool {
		if cUse[j] || slashUse[i+j] || backSlashUse[i-j+n-1] {
			return false
		}
		return true
	}
	var dfs func(r, step int)
	dfs = func(r, step int) {
		if step == n {
			count++
			return
		}
		for c := 0; c < n; c++ {
			if !isValid(r, c) {
				continue
			}
			solve[r][c] = 'Q'
			setUse(r, c)
			dfs(r+1, step+1)
			clearUse(r, c)
			solve[r][c] = '.'
		}
	}
	init()
	dfs(0, 0)

	return count
}
