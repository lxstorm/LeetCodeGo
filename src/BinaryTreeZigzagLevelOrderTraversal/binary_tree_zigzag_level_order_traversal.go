package BinaryTreeZigzagLevelOrderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// use count method
func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	queue := []*TreeNode{}
	oddLevel := true
	if root == nil {
		return ret
	}
	queue = append(queue, root)
	for len(queue) != 0 {
		count := len(queue)
		tmp := make([]int, count)
		for i := 0; i < count; i++ {
			front := queue[0]
			queue = queue[1:]
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
			if oddLevel {
				tmp[i] = front.Val
			} else {
				tmp[count-1-i] = front.Val
			}
		}
		ret = append(ret, tmp)
		oddLevel = !oddLevel
	}
	return ret
}
