package RemoveNthNodeFromEndOfList

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	l := dummy
	r := head
	for i := 0; i < n; i++ {
		r = r.Next
	}
	for ; r != nil; l, r = l.Next, r.Next {
	}
	l.Next = l.Next.Next

	return dummy.Next
}
