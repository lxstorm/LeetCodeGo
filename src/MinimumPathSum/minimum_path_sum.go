package MinimumPathSum

import (
	"math"
)

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

// can also pre allocate value in border, not to create a r+1, c+1 matrix
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	r, c := len(grid), len(grid[0])
	mem := [][]int{}
	for i := 0; i < r+1; i++ {
		tmp := make([]int, c+1)
		for j := 0; j < c+1; j++ {
			tmp[j] = math.MaxInt32
		}
		mem = append(mem, tmp)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i == 0 && j == 0 {
				mem[i+1][j+1] = grid[i][j]
			} else {
				mem[i+1][j+1] = grid[i][j] + min(mem[i][j+1], mem[i+1][j])
			}
		}
	}

	return mem[r][c]
}
