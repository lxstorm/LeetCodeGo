package RotateImage

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
	rotate(input)
	for i := 0; i < len(input); i++ {
		fmt.Println(input[i])
	}
}
