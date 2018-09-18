package UniquePathII

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	input := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	result := uniquePathsWithObstacles(input)
	fmt.Println(result)
}
