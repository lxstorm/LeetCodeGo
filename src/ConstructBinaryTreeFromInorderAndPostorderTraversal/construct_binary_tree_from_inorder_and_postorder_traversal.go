package ConstructBinaryTreeFromInorderAndPostorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	rootVal := postorder[len(postorder)-1]
	root.Val = rootVal
	rootIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			rootIndex = i
			break
		}
	}
	root.Left = buildTree(inorder[0:rootIndex], postorder[0:rootIndex])
	root.Right = buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])

	return root
}
