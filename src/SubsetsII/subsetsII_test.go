package SubsetsII

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	input := []int{1, 2, 2}
	result := subsetsWithDup(input)
	fmt.Println(result)
}
