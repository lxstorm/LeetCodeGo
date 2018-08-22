package PalindromeNumber

import "testing"

type testCase struct {
	n   int
	ans bool
}

func TestPalinNum(t *testing.T) {
	testCases := []testCase{
		testCase{
			n:   121,
			ans: true,
		},
		testCase{
			n:   -121,
			ans: false,
		},
		testCase{
			n:   10,
			ans: false,
		},
	}
	for _, curCase := range testCases {
		result := isPalindrome(curCase.n)
		if result != curCase.ans {
			t.Errorf("Input: %v, want: %v, result: %v",
				curCase.n,
				curCase.ans,
				result)
		}
	}

}
