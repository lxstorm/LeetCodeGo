package UniquePath

import "testing"

type testCase struct {
	m   int
	n   int
	ans int
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{
			m:   3,
			n:   2,
			ans: 3,
		},
		testCase{
			m:   7,
			n:   3,
			ans: 28,
		},
	}
	for _, curCase := range testCases {
		result := uniquePaths(curCase.m, curCase.n)
		if result != curCase.ans {
			t.Errorf("Test case is %v, result is %v",
				curCase,
				result)
		}
	}

}
