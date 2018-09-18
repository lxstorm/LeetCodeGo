package SimplifyPath

import "testing"

type testCase struct {
	path string
	ans  string
}

func TestF(t *testing.T) {
	testCases := []testCase{
		testCase{"/home/", "/home"},
		testCase{"/a/./b/../../c/", "/c"},
		testCase{"/a/./b/../../c/", "/c"},
		testCase{"/a/../../b/../c//.//", "/c"},
		testCase{"/a//b////c/d//././/..", "/a/b/c"},
	}
	for _, curCase := range testCases {
		result := simplifyPath(curCase.path)
		if result != curCase.ans {
			t.Errorf("Test case is %v, result is %v", curCase, result)
		}
	}
}
