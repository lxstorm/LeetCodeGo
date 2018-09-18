package Combinations

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	n, k := 4, 2
	output := combine(n, k)
	fmt.Println(output)
}
