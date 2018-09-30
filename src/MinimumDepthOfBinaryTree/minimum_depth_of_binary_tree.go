package MinimumDepthOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

// recursive solution
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

// BFS solution
func minDepthBFS(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) != 0 {
		count := len(queue)
		for i := 0; i < count; i++ {
			front := queue[0]
			queue = queue[1:]
			if front.Left == nil && front.Right == nil {
				return depth + 1
			}
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		}
		depth++
	}
	return depth
}
