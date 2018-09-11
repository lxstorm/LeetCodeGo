package SwapNodesInPairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var p *ListNode
	p = head.Next
	head.Next = swapPairs(p.Next)
	p.Next = head

	return p
}
