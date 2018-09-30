package BinaryTreeLevelOrderTraversalII

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) != 0 {
		count := len(queue)
		tmp := []int{}
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
		ret = append([][]int{tmp}, ret...)
	}

	return ret
}
