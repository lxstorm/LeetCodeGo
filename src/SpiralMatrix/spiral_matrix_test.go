package SpiralMatrix

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	input := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	result := spiralOrder(input)
	fmt.Println(result)
}
