package AddTwoSum

// ListNode is user-defined struct
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertNode(l *ListNode, val int) *ListNode {
	if l == nil {
		return nil
	}
	newNode := new(ListNode)
	newNode.Val = val
	newNode.Next = l.Next
	l.Next = newNode

	return newNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	pResult := result
	var p1Val, p2Val int
	carry, remainder := 0, 0
	// add carry not eq zero condition in for cond,
	// avoid extra add node after for if carry is not equal 0
	for p1, p2 := l1, l2; p1 != nil || p2 != nil || carry != 0; {
		p1Val, p2Val = 0, 0
		if p1 != nil {
			p1Val = p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			p2Val = p2.Val
			p2 = p2.Next
		}
		sum := p1Val + p2Val + carry
		carry = sum / 10
		remainder = sum % 10
		pResult = insertNode(pResult, remainder)
	}

	return result.Next
}
