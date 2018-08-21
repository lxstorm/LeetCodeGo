package ZigZagConversion

import (
	"fmt"
	"testing"
)

type testCase struct {
	s     string
	nRows int
	ans   string
}

func TestZigZag(t *testing.T) {
	testCases := []testCase{
		testCase{
			s:     "PAYPALISHIRING",
			nRows: 3,
			ans:   "PAHNAPLSIIGYIR",
		},
		testCase{
			s:     "PAYPALISHIRING",
			nRows: 4,
			ans:   "PINALSIGYAHRPI",
		},
		testCase{
			s:     "AB",
			nRows: 1,
			ans:   "AB",
		},
	}
	for _, curCase := range testCases {
		result := convert(curCase.s, curCase.nRows)
		if result != curCase.ans {
			t.Errorf("input: %v, rows: %v, want: %v, calculated: %v",
				curCase.s,
				curCase.nRows,
				curCase.ans,
				result,
			)
			fmt.Println(len(curCase.ans), len(result))
			fmt.Println([]byte(result))
		}
	}

}
