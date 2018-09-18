package AddBinary

import "testing"

type testCase struct {
	a   string
	b   string
	ans string
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{a: "11", b: "1", ans: "100"},
		testCase{a: "1010", b: "1011", ans: "10101"},
	}
	for _, curCase := range testCases {
		result := addBinary(curCase.a, curCase.b)
		if result != curCase.ans {
			t.Errorf("Test case is %v, result is %v", curCase, result)
		}
	}
}
