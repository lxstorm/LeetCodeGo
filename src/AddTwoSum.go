/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res := new(ListNode)
    res.Next = nil
    p1, p2 := l1, l2
    cur := res
    carry := 0

    for ; p1 != nil || p2 != nil; {
        tmpNode := new(ListNode)
        var v1, v2 int

        if p1 != nil {
            v1 = p1.Val
        }
        if p2 != nil {
            v2 = p2.Val
        }
        sum := v1 + v2 + carry
        tmpNode.Val = sum % 10
        tmpNode.Next = nil
        carry = sum / 10

        cur.Next = tmpNode
        cur = tmpNode

        if p1 != nil {
            p1 = p1.Next
        }
        if p2 != nil {
            p2 = p2.Next
        }
    }

    if carry != 0 {
        tmpNode := new(ListNode)
        tmpNode.Val = 1
        tmpNode.Next = nil
        cur.Next = tmpNode
    }

    return res.Next

}
