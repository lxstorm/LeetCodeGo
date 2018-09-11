package DivideTwoIntegers

import "testing"

type testCase struct {
	dividen int
	dividor int
	ans     int
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{
			dividen: 10,
			dividor: 3,
			ans:     3,
		},
		testCase{
			dividen: 7,
			dividor: -3,
			ans:     -2,
		},
	}
	for _, curCase := range testCases {
		result := divide(curCase.dividen, curCase.dividor)
		if result != curCase.ans {
			t.Errorf("test case is %v, result is %v",
				curCase,
				result)

		}
	}

}
