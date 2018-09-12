package SpiralMatrix

type direction struct {
	x int
	y int
}

type position struct {
	x int
	y int
}

func spiralOrder(matrix [][]int) []int {
	ret := []int{}
	if len(matrix) == 0 {
		return ret
	}
	r, c := len(matrix), len(matrix[0])
	directions := []direction{
		direction{0, 1},
		direction{1, 0},
		direction{0, -1},
		direction{-1, 0},
	}
	m := [][]bool{}
	for i := 0; i < r; i++ {
		m = append(m, make([]bool, c))
	}
	isValidPos := func(p position) bool {
		i, j := p.x, p.y
		if i < 0 || i >= r || j < 0 || j >= c || m[i][j] {
			return false
		}
		return true
	}
	getNextPos := func(p position, d direction) position {
		return position{p.x + d.x, p.y + d.y}
	}
	dirCnt := 0
	curPos := position{0, 0}
	for i := 0; i < r*c; i++ {
		ret = append(ret, matrix[curPos.x][curPos.y])
		m[curPos.x][curPos.y] = true
		nextPos := getNextPos(curPos, directions[dirCnt%4])
		if !isValidPos(nextPos) {
			dirCnt++
			curPos = getNextPos(curPos, directions[dirCnt%4])
		} else {
			curPos = nextPos
		}
	}

	return ret
}
