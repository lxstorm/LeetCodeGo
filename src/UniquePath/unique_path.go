package UniquePath

func uniquePaths(m int, n int) int {
	mem := [][]int{}
	for i := 0; i < n; i++ {
		mem = append(mem, make([]int, m))
	}
	for i := 0; i < m; i++ {
		mem[0][i] = 1
	}
	for i := 0; i < n; i++ {
		mem[i][0] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			mem[i][j] = mem[i][j-1] + mem[i-1][j]
		}
	}

	return mem[n-1][m-1]
}
