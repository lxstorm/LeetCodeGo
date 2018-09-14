package RotateList

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	p, n := head, 1
	for ; p.Next != nil; p, n = p.Next, n+1 {
	}
	p.Next = head
	k = n - k%n
	// now p at the last element, try to stop p at new start
	// p stop at new start
	start := head
	for i := 0; i < k; i++ {
		p = p.Next
		start = start.Next
	}
	p.Next = nil

	return start
}
