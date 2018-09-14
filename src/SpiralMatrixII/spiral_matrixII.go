package SpiralMatrixII

type dircection struct {
	x int
	y int
}

type position struct {
	x int
	y int
}

func generateMatrix(n int) [][]int {
	matrix := [][]int{}
	for i := 0; i < n; i++ {
		matrix = append(matrix, make([]int, n))
	}

	isValidPos := func(p position) bool {
		if p.x >= n || p.x < 0 || p.y >= n || p.y < 0 || matrix[p.x][p.y] != 0 {
			return false
		}
		return true
	}
	getNextPos := func(p position, d dircection) position {
		return position{p.x + d.x, p.y + d.y}
	}

	directions := []dircection{
		dircection{0, 1},
		dircection{1, 0},
		dircection{0, -1},
		dircection{-1, 0},
	}
	directCnt := 0
	curPos := position{0, 0}

	for i := 0; i < n*n; i++ {
		matrix[curPos.x][curPos.y] = i + 1
		curDirect := directions[directCnt%4]
		nextPos := getNextPos(curPos, curDirect)
		if !isValidPos(nextPos) {
			directCnt++
			curDirect = directions[directCnt%4]
			nextPos = getNextPos(curPos, curDirect)
		}
		curPos = nextPos
	}

	return matrix
}
