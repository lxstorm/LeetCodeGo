package LetterCombinationsOfAPhoneNumber

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	digits := "23"
	result := letterCombinations(digits)
	fmt.Printf("Given: %v, Result: %v\n",
		digits,
		result)
}
