package ValidateBinarySearchTree

import (
	"fmt"
	"testing"
)

func cloneTree(t *TreeNode) *TreeNode {
	root := new(TreeNode)
	left := cloneTree(t.Left)
	right := cloneTree(t.Right)
	root.Left = left
	root.Right = right

	return root
}

func TestF(t *testing.T) {
	t1 := TreeNode{1, nil, nil}
	t2 := TreeNode{2, nil, nil}
	t3 := TreeNode{3, nil, nil}
	t4 := TreeNode{4, nil, nil}
	// t5 := TreeNode{5, nil, nil}
	t2.Left = &t1
	t2.Right = &t3
	t1.Right = &t4
	res := isValidBST(&t2)
	fmt.Println(res)
	res = isValidBSTInOrder(&t2)
	fmt.Println(res)
}
