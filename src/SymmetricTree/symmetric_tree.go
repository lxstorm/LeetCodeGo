package SymmetricTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs solution
func isSymmetric(root *TreeNode) bool {
	stack := []*TreeNode{}
	stack = append(stack, root)
	stack = append(stack, root)
	for len(stack) != 0 {
		leftTop := stack[len(stack)-2]
		rightTop := stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		if leftTop == nil && rightTop == nil {
			continue
		}
		if leftTop == nil || rightTop == nil || leftTop.Val != rightTop.Val {
			return false
		}
		stack = append(stack, leftTop.Left)
		stack = append(stack, rightTop.Right)
		stack = append(stack, leftTop.Right)
		stack = append(stack, rightTop.Left)
	}
	return true
}

// recursive solution
func isSymmetricRecursive(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var helper func(*TreeNode, *TreeNode) bool
	helper = func(t1, t2 *TreeNode) bool {
		if t1 == nil && t2 == nil {
			return true
		}
		if t1 == nil || t2 == nil {
			return false
		}
		return t1.Val == t2.Val && helper(t1.Left, t2.Right) && helper(t1.Right, t2.Left)
	}

	return helper(root.Left, root.Right)
}
