package Powxn

import "testing"

type testCase struct {
	x   float64
	n   int
	ans float64
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{
			x:   2.00000,
			n:   10,
			ans: 1024.0000,
		},
		testCase{
			x:   2.10000,
			n:   3,
			ans: 9.26100,
		},
		testCase{
			x:   2.00000,
			n:   -2,
			ans: 0.25000,
		},
	}
	for _, curCase := range testCases {
		result := myPow(curCase.x, curCase.n)
		if result != curCase.ans {
			t.Errorf("Test case is %v, result is %v",
				curCase,
				result)
		}
	}
}
