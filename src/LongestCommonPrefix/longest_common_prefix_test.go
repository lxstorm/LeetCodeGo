package LongestCommonPrefix

import "testing"

type testCase struct {
	input []string
	ans   string
}

func TestPrefix(t *testing.T) {
	testCases := []testCase{
		testCase{
			input: []string{"flower", "flow", "flight"},
			ans:   "fl",
		},
		testCase{
			input: []string{"dog", "racecar", "car"},
			ans:   "",
		},
		testCase{
			input: []string{"aa", "a"},
			ans:   "a",
		},
	}
	for _, curCase := range testCases {
		result := longestCommonPrefix(curCase.input)
		if result != curCase.ans {
			t.Errorf("want: %v, result: %v",
				curCase.ans,
				result)
		}
	}
}
