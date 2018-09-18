package PlusOne

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	input := []int{1, 2, 3}
	// input := []int{9, 9, 9}
	output := plusOne(input)
	fmt.Println(output)
}
