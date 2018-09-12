package NQueens

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	n := 8
	result := solveNQueens(n)
	for i := 0; i < len(result); i++ {
		for j := 0; j < n; j++ {
			fmt.Println(result[i][j])
		}
		fmt.Println("===========")
	}
}
