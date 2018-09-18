package WordSearch

import (
	"testing"
)

type testCase struct {
	board [][]byte
	word  string
	ans   bool
}

func TestF(t *testing.T) {
	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	testCases := []testCase{
		testCase{
			board: board,
			word:  "ABCCED",
			ans:   true,
		},
		testCase{
			board: board,
			word:  "SEE",
			ans:   true,
		},
		testCase{
			board: board,
			word:  "ABCB",
			ans:   false,
		},
	}
	for _, curCase := range testCases {
		result := exist(curCase.board, curCase.word)
		if result != curCase.ans {
			t.Errorf("Test case is %v, result is %v", curCase, result)
		}
	}
}
