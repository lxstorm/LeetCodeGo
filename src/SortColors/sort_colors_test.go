package SortColors

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	input := []int{2, 0, 2, 1, 1, 0}
	fmt.Println(input)
	sortColors(input)
	fmt.Println(input)
}
