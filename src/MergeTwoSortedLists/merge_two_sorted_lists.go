package MergeTwoSortedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	p1, p2, r := l1, l2, dummy
	for p1 != nil && p2 != nil {
		tmp := p1
		if p1.Val > p2.Val {
			tmp = p2
			p2 = p2.Next
		} else {
			p1 = p1.Next
		}
		r.Next = tmp
		r.Next.Next = nil
		r = r.Next
	}
	if p1 != nil {
		r.Next = p1
	}
	if p2 != nil {
		r.Next = p2
	}
	return dummy.Next
}
