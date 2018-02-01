/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry, val1, val2, sum, remainder := 0, 0, 0, 0, 0
	p1, p2 := l1, l2
	res := ListNode{Val: 0, Next: nil}
	p3 := &res

	for p1 != nil || p2 != nil || carry != 0 {
		node := new(ListNode)
		val1, val2 = 0, 0
		if p1 != nil {
			val1 = p1.Val
		}
		if p2 != nil {
			val2 = p2.Val
		}
		sum = val1 + val2 + carry
		carry = sum / 10
		remainder = sum % 10

		node.Val = remainder
		node.Next = nil
		p3.Next = node
		p3 = node
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}
	return res.Next
}

