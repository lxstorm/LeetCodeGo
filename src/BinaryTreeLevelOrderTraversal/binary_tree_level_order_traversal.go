package BinaryTreeLevelOrderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// mark method
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	tmp := []int{}
	if root == nil {
		return ret
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	queue = append(queue, nil)

	for len(queue) != 0 {
		front := queue[0]
		queue = queue[1:]
		if front == nil {
			ret = append(ret, append([]int(nil), tmp...))
			tmp = []int{}
			if len(queue) != 0 {
				queue = append(queue, nil)
			}
			continue
		}
		tmp = append(tmp, front.Val)
		if front.Left != nil {
			queue = append(queue, front.Left)
		}
		if front.Right != nil {
			queue = append(queue, front.Right)
		}

	}
	return ret
}

// count method
func levelOrderUseCounter(root *TreeNode) [][]int {
	ret := [][]int{}
	tmp := []int{}
	if root == nil {
		return ret
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) != 0 {
		count := len(queue)
		for i := 0; i < count; i++ {
			front := queue[0]
			queue = queue[1:]
			tmp = append(tmp, front.Val)
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		}
		ret = append(ret, append([]int(nil), tmp...))
		tmp = []int{}
	}

	return ret
}
