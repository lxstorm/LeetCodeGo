package PathSumII

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	t1 := &TreeNode{3, nil, nil}
	t2 := &TreeNode{9, nil, nil}
	t3 := &TreeNode{20, nil, nil}
	t4 := &TreeNode{15, nil, nil}
	t5 := &TreeNode{7, nil, nil}
	t1.Left = t2
	t1.Right = t3
	t3.Left = t4
	t3.Right = t5

	result := pathSum(t1, 30)
	fmt.Println(result)
	result = pathSumDFS(t1, 30)
	fmt.Println(result)
}
