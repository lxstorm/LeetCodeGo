package BinaryTreeInorderTraversal

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	t1 := TreeNode{1, nil, nil}
	t2 := TreeNode{2, nil, nil}
	t3 := TreeNode{3, nil, nil}
	t1.Right = &t2
	t2.Left = &t3
	result := inorderTraversal(&t1)
	fmt.Println(result)
}
