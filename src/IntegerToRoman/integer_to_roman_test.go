package IntegerToRoman

import "testing"

type testCase struct {
	i   int
	ans string
}

func TestConvert(t *testing.T) {
	testCases := []testCase{
		testCase{
			i:   3,
			ans: "III",
		},
		testCase{
			i:   4,
			ans: "IV",
		},
		testCase{
			i:   9,
			ans: "IX",
		},
		testCase{
			i:   58,
			ans: "LVIII",
		},
		testCase{
			i:   1994,
			ans: "MCMXCIV",
		},
	}
	for _, curCase := range testCases {
		result := intToRoman(curCase.i)
		if result != curCase.ans {
			t.Errorf("Test case %v, result: %v",
				curCase,
				result)
		}
	}
}
