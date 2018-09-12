package MultiplyStrings

import "testing"

type testCase struct {
	s1  string
	s2  string
	ans string
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{
			s1:  "123",
			s2:  "456",
			ans: "56088",
		},
		testCase{
			s1:  "2",
			s2:  "3",
			ans: "6",
		},
		testCase{
			s1:  "999",
			s2:  "999",
			ans: "998001",
		},
	}
	for _, curCase := range testCases {
		result := multiply(curCase.s1, curCase.s2)
		if result != curCase.ans {
			t.Errorf("test case is: %v, result is %v",
				curCase,
				result)
		}
	}
}
