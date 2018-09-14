package MergeIntevals

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	// input := []Interval{
	// 	Interval{1, 3},
	// 	Interval{2, 6},
	// 	Interval{8, 10},
	// 	Interval{15, 18},
	// }
	input := []Interval{
		Interval{1, 4},
		Interval{4, 5},
	}
	result := merge(input)
	fmt.Println(result)
}
