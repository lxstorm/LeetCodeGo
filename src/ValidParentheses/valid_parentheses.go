package ValidParentheses

func isValid(s string) bool {
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	pushStack := func(c byte) {
		stack = append(stack, c)
	}
	popStack := func() {
		if len(stack) == 0 {
			return
		}
		stack = stack[0 : len(stack)-1]
	}
	var topElement byte
	for _, c := range []byte(s) {
		if len(stack) == 0 {
			pushStack(c)
			continue
		}
		topElement = stack[len(stack)-1]
		mVal, ok := m[c]
		if ok && topElement == mVal {
			popStack()
			continue
		}
		pushStack(c)
	}

	return len(stack) == 0
}
