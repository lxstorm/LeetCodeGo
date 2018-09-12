package NQueensII

import "testing"

type testCase struct {
	n   int
	ans int
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{
			n:   4,
			ans: 2,
		},
	}
	for _, curCase := range testCases {
		result := totalNQueens(curCase.n)
		if result != curCase.ans {
			t.Errorf("Test case is %v, result is %v",
				curCase,
				result)
		}
	}
}
