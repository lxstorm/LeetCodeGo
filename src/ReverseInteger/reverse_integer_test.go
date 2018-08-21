package ReverseInteger

import "testing"

type testCase struct {
	i   int
	ans int
}

func TestReverse(t *testing.T) {
	testCases := []testCase{
		testCase{
			i:   123,
			ans: 321,
		},
		testCase{
			i:   -321,
			ans: -123,
		},
		testCase{
			i:   1534236469,
			ans: 0,
		},
	}
	for _, curCase := range testCases {
		result := reverse(curCase.i)
		if result != curCase.ans {
			t.Errorf("input: %v, want: %v, result: %v",
				curCase.i,
				curCase.ans,
				result)
		}
	}
}
