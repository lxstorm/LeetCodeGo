package ImplementstrStr

import "testing"

type testCase struct {
	haystack string
	needle   string
	ans      int
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{
			haystack: "hello",
			needle:   "ll",
			ans:      2,
		},
		testCase{
			haystack: "aaaaa",
			needle:   "bba",
			ans:      -1,
		},
		testCase{
			haystack: "abbbaaaaaaabbababbbbabababbbbbbbaaaaaaabbaaabbaababbbbababababaabbbbbbaaaaababbbbaaabababbbaaaabbbaabbbbbbabababbabaaaaabaabaaababbbaaabaababbaaabaaababbabbbbababaaaaaaababaabaabbaabbbaaabaaaaaa",
			needle:   "aabaaaabababbbabababbbaabaabaaaaabaabbbaabbbbba",
			ans:      -1,
		},
	}
	for _, curCase := range testCases {
		result := strStr(curCase.haystack, curCase.needle)
		if result != curCase.ans {
			t.Errorf("test case is %v, result is %v",
				curCase,
				result)
		}
	}

}
