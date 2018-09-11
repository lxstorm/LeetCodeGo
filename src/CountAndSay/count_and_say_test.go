package CountAndSay

import "testing"

type testCase struct {
	n   int
	ans string
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{
			1,
			"1",
		},
		testCase{
			2,
			"11",
		},
		testCase{
			3,
			"21",
		},
		testCase{
			4,
			"1211",
		},
	}
	for _, curCase := range testCases {
		result := countAndSay(curCase.n)
		if result != curCase.ans {
			t.Errorf("test case int is %v, result is %v",
				curCase.n,
				result)
		}

	}

}
