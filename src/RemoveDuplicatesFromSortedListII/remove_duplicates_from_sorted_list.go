package RemoveDuplicatesFromSortedListII

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := new(ListNode)
	p := dummy
	slow := head
	for slow != nil {
		fast := slow.Next
		for ; fast != nil && fast.Val == slow.Val; fast = fast.Next {
		}
		if fast == slow.Next {
			p.Next = slow
			p = slow
			p.Next = nil
		}
		slow = fast
	}

	return dummy.Next

}
