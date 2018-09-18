package SubSets

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	nums := []int{1, 2, 3}
	// result := subsets(nums)
	result := subsetsBacktrace(nums)
	fmt.Println(result)
}
