package RemoveDuplicatesFromSortedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	for slow != nil {
		fast := slow.Next
		for ; fast != nil && fast.Val == slow.Val; fast = fast.Next {
		}
		slow.Next = fast
		slow = fast
	}

	return head
}
