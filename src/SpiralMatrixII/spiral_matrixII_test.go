package SpiralMatrixII

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	input := 3
	result := generateMatrix(input)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
