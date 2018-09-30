package ReverseNodesInKGroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := new(ListNode)
	dummy.Next = head
	count := 0
	for pt := head; pt != nil; pt = pt.Next {
		count++
	}
	if count < k {
		return head
	}

	p := head
	for i := 0; p != nil && i < k; i++ {
		tmp := p.Next
		p.Next = dummy.Next
		dummy.Next = p
		p = tmp
	}
	head.Next = reverseKGroup(p, k)

	return dummy.Next
}
