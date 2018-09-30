package ReverseLinkedListII

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	count := 0
	dummyOri := new(ListNode)
	dummyOri.Next = head
	pushFromNode := dummyOri
	pushTailNode := dummyOri
	// no need to complete in a full loop
	for p := dummyOri; p != nil && p.Next != nil; {
		// p.Next is the count-th node
		count++
		curNode := p.Next
		if count == m {
			pushFromNode = p
			pushTailNode = curNode
		}
		if count > m && count <= n {
			pushTailNode.Next = curNode.Next
			curNode.Next = pushFromNode.Next
			pushFromNode.Next = curNode
		}
		if count <= m || count >= n {
			p = p.Next
		}
	}
	return dummyOri.Next
}

func reverseBetweenOptimized(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	dummyOri := new(ListNode)
	dummyOri.Next = head
	pre := dummyOri
	count := n - m
	for ; m > 1; m-- {
		pre = pre.Next
	}
	start, then := pre.Next, pre.Next.Next
	for ; count > 0; count-- {
		start.Next = then.Next
		then.Next = pre.Next
		pre.Next = then
		then = start.Next
	}
	return dummyOri.Next
}
