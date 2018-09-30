package ConstructBinaryTreeFromPreorderAndInorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	rootVal := preorder[0]
	root.Val = rootVal
	rootIndex := 0
	for ; inorder[rootIndex] != rootVal; rootIndex++ {
	}
	leftTreeLen := rootIndex
	root.Left = buildTree(preorder[1:leftTreeLen+1], inorder[0:leftTreeLen])
	root.Right = buildTree(preorder[leftTreeLen+1:], inorder[leftTreeLen+1:])

	return root
}
