package PermutationSequence

import "testing"

type testCase struct {
	n   int
	k   int
	ans string
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{3, 3, "213"},
		testCase{4, 9, "2314"},
	}
	for _, curCase := range testCases {
		result := getPermutation(curCase.n, curCase.k)
		if result != curCase.ans {
			t.Errorf("Testcase is %v, result is %v",
				curCase,
				result)
		}
	}

}
