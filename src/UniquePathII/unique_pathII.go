package UniquePathII

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	mem := [][]int{}
	r, c := len(obstacleGrid), len(obstacleGrid[0])
	for i := 0; i < r+1; i++ {
		mem = append(mem, make([]int, c+1))
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if obstacleGrid[i][j] == 1 {
				mem[i+1][j+1] = 0
			} else {
				if i == 0 && j == 0 {
					mem[i+1][j+1] = 1
				} else {
					mem[i+1][j+1] = mem[i][j+1] + mem[i+1][j]
				}
			}
		}
	}

	return mem[r][c]

}
