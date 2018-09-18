package PartitionList

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	lessDummy, notLessDummy := new(ListNode), new(ListNode)
	pLess, pNotLess := lessDummy, notLessDummy
	cur := head
	for cur != nil {
		if cur.Val < x {
			pLess.Next = cur
			cur = cur.Next
			pLess = pLess.Next
			pLess.Next = nil
		} else {
			pNotLess.Next = cur
			cur = cur.Next
			pNotLess = pNotLess.Next
			pNotLess.Next = nil
		}
	}
	pLess.Next = notLessDummy.Next

	return lessDummy.Next
}
