package PathSumII

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive solution
func pathSum(root *TreeNode, sum int) [][]int {
	ret := [][]int{}
	path := []int{}

	var helper func(*TreeNode, int)
	helper = func(curNode *TreeNode, target int) {
		if curNode == nil {
			return
		}
		if curNode.Left == nil && curNode.Right == nil {
			if curNode.Val == target {
				path = append(path, curNode.Val)
				ret = append(ret, append([]int(nil), path...))
				path = path[:len(path)-1]
			}
			return
		}
		if curNode.Left != nil {
			path = append(path, curNode.Val)
			helper(curNode.Left, target-curNode.Val)
			path = path[:len(path)-1]
		}
		if curNode.Right != nil {
			path = append(path, curNode.Val)
			helper(curNode.Right, target-curNode.Val)
			path = path[:len(path)-1]
		}
	}

	helper(root, sum)

	return ret
}

// iterative solution
func pathSumDFS(root *TreeNode, sum int) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	stack := []*TreeNode{}
	pathStack := [][]int{}
	stack = append(stack, root)
	pathStack = append(pathStack, []int{root.Val})
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		tmpPath := pathStack[len(pathStack)-1]
		pathStack = pathStack[0 : len(pathStack)-1]

		if top.Left == nil && top.Right == nil {
			s := 0
			for i := 0; i < len(tmpPath); i++ {
				s += tmpPath[i]
			}
			if s == sum {
				ret = append(ret, tmpPath)
			}
		}
		if top.Right != nil {
			stack = append(stack, top.Right)
			newPath := append([]int(nil), tmpPath...)
			newPath = append(newPath, top.Right.Val)
			pathStack = append(pathStack, newPath)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
			newPath := append([]int(nil), tmpPath...)
			newPath = append(newPath, top.Left.Val)
			pathStack = append(pathStack, newPath)
		}
	}

	return ret
}
