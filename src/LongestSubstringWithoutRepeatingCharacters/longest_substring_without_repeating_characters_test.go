package LongestSubstringWithoutRepeatingCharacters

import "testing"

type testCase struct {
	s      string
	substr string
	maxLen int
}

func TestLongestSubstring(t *testing.T) {
	testCases := []testCase{
		testCase{
			s:      "abcabcbb",
			substr: "abc",
			maxLen: 3,
		},
		testCase{
			s:      "bbbbb",
			substr: "b",
			maxLen: 1,
		},
		testCase{
			s:      "pwwkew",
			substr: "wke",
			maxLen: 3,
		},
		testCase{
			s:      "p",
			substr: "p",
			maxLen: 1,
		},
	}
	for _, val := range testCases {
		result := lengthOfLongestSubstring(val.s)
		if result != val.maxLen {
			t.Errorf("Max len not match: want %v, caculated %v",
				val.maxLen,
				result,
			)
		}
	}
}
