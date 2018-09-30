package SymmetricTree

import "testing"

func TestF(t *testing.T) {
	t1 := &TreeNode{1, nil, nil}
	t2 := &TreeNode{2, nil, nil}
	t3 := &TreeNode{2, nil, nil}
	t4 := &TreeNode{3, nil, nil}
	t5 := &TreeNode{4, nil, nil}
	t6 := &TreeNode{4, nil, nil}
	t7 := &TreeNode{3, nil, nil}

	t1.Left = t2
	t1.Right = t3
	t2.Left = t4
	t2.Right = t5
	t3.Left = t6
	t3.Right = t7

	result := isSymmetric(t1)
	if result != true {
		t.Errorf("NOT PASS")
	}
	result = isSymmetricRecursive(t1)
	if result != true {
		t.Errorf("NOT PASS")
	}
}
