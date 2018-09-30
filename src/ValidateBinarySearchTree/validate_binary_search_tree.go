package ValidateBinarySearchTree

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var helper func(t *TreeNode, upLimit int, lowLimit int, upFlag bool, lowFlag bool) bool
	helper = func(t *TreeNode, upLimit int, lowLimit int, upFlag bool, lowFlag bool) bool {
		if t == nil {
			return true
		}
		if (upFlag && t.Val >= upLimit) || (lowFlag && t.Val <= lowLimit) {
			return false
		}
		return helper(t.Left, t.Val, lowLimit, true, lowFlag) &&
			helper(t.Right, upLimit, t.Val, upFlag, true)
	}

	return helper(root, math.MaxInt32, math.MinInt32, false, false)
}

func isValidBSTInOrder(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := []*TreeNode{}
	p := root
	pre := (*TreeNode)(nil)
	for p != nil || len(stack) != 0 {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		} else {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pre != nil && pre.Val >= p.Val {
				return false
			}
			pre = p
			p = p.Right
		}
	}
	return true
}
