package SetMatrixZeroes

import (
	"fmt"
	"testing"
)

func printMaxtrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}

func TestF(t *testing.T) {
	// input := [][]int{
	// 	[]int{1, 1, 1},
	// 	[]int{1, 0, 1},
	// 	[]int{1, 1, 1},
	// }
	input := [][]int{
		[]int{0, 1, 2, 0},
		[]int{3, 4, 5, 2},
		[]int{1, 3, 1, 5},
	}
	printMaxtrix(input)
	fmt.Println("===========")
	// setZeroes(input)
	setZeroesConstant(input)
	printMaxtrix(input)
}
