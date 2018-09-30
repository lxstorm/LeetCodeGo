package FlattenBinaryTreeToLinkedList

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	dummyLeft := new(TreeNode)
	dummyLeft.Right = root.Left
	cur := dummyLeft
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = root.Right
	root.Left = nil
	root.Right = dummyLeft.Right

}
