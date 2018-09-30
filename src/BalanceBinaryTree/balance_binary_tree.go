package BalanceBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	ldepth, rdepth := depth(root.Left), depth(root.Right)
	if ldepth-rdepth > 1 || rdepth-ldepth > 1 {
		return false
	}
	if isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}
	return false
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ldepth, rdepth := depth(root.Left), depth(root.Right)
	if ldepth > rdepth {
		return ldepth + 1
	}
	return rdepth + 1
}
