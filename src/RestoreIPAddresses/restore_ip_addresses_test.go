package RestoreIPAddresses

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	// input := "25525511135"
	input := "010010"
	output := restoreIpAddresses(input)
	fmt.Println(output)
}
