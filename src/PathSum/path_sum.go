package PathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, sum-root.Val) ||
		hasPathSum(root.Right, sum-root.Val)
}

// just add the sum for every node, no need to perform add and substract
func hasPathSumDFS(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	curVal := []int{}
	curVal = append(curVal, root.Val)
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		val := curVal[len(curVal)-1]
		curVal = curVal[0 : len(curVal)-1]
		if top.Right == nil && top.Left == nil {
			if val == sum {
				return true
			}
		}
		if top.Right != nil {
			stack = append(stack, top.Right)
			curVal = append(curVal, top.Right.Val+val)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
			curVal = append(curVal, top.Left.Val+val)
		}
	}
	return false
}
