package GenerateParentheses

func generateParenthesis(n int) []string {
	ret := []string{}
	var helper func(string, int, int)
	helper = func(s string, left int, right int) {
		if left == 0 && right == 0 {
			ret = append(ret, s)
		}
		if left > 0 {
			helper(s+"(", left-1, right+1)
		}
		if right > 0 {
			helper(s+")", left, right-1)
		}
	}
	helper("", n, 0)
	return ret

}
