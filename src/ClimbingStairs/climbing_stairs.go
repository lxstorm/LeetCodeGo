package ClimbingStairs

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	preTwoJump := 1
	preOneJump := 2

	for i := 3; i <= n; i++ {
		solution := preOneJump + preTwoJump
		preTwoJump = preOneJump
		preOneJump = solution
	}

	return preOneJump

}
