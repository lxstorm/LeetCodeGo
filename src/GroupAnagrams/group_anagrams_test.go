package GroupAnagrams

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)
	fmt.Println(result)
}
