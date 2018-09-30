package ConvertSortedListToBinarySearchTree

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	return helper(head, nil)
}

func helper(head *ListNode, tail *ListNode) *TreeNode {
	// invariant treenode value [head, tail)
	if head == tail {
		return nil
	}
	slow, fast := head, head
	for fast.Next != tail && fast.Next.Next != tail {
		slow = slow.Next
		fast = fast.Next.Next
	}
	root := new(TreeNode)
	root.Val = slow.Val
	root.Left = helper(head, slow)
	root.Right = helper(slow.Next, tail)

	return root
}
