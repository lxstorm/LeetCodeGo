package MaximumDepthOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// can also use recursive dfs
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	levelCount := 0
	queue = append(queue, root)
	for len(queue) != 0 {
		count := len(queue)
		for i := 0; i < count; i++ {
			front := queue[0]
			queue = queue[1:]
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		}
		levelCount++
	}
	return levelCount
}
