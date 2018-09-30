package SameTree

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	t1 := TreeNode{1, nil, nil}
	t2 := TreeNode{2, nil, nil}
	t3 := TreeNode{3, nil, nil}
	t2.Left = &t1
	t2.Right = &t3

	t4 := TreeNode{1, nil, nil}
	t5 := TreeNode{2, nil, nil}
	t6 := TreeNode{3, nil, nil}
	t5.Left = &t4
	t5.Right = &t6

	result := isSameTree(&t2, &t5)
	fmt.Println(result)
	result = isSameTreeDFS(&t2, &t5)
	fmt.Println(result)
}
