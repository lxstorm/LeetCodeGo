package UniqueBinarySearchTree

func numTrees(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	mem := make([]int, n+1)
	mem[0] = 1
	mem[1] = 1
	for i := 2; i <= n; i++ {
		res := 0
		for j := 0; j < i; j++ {
			// left has j elements, right has i-1-j elements
			res += mem[j] * mem[i-1-j]
		}
		mem[i] = res
	}

	return mem[n]
}
