package AddTwoSum

import (
	"strconv"
	"strings"
	"testing"
)

type TestCase struct {
	op1 *ListNode
	op2 *ListNode
	ans *ListNode
}

func generateListNode(op []int) *ListNode {
	if op == nil {
		return nil
	}
	dummy := new(ListNode)
	dummy.Next = nil
	curNode := dummy
	for _, value := range op {
		newNode := new(ListNode)
		newNode.Val = value
		newNode.Next = curNode.Next
		curNode.Next = newNode

		curNode = curNode.Next
	}

	return dummy.Next
}

func prettyPrintList(l *ListNode) string {
	s := []string{}
	var output strings.Builder
	for p := l; p != nil; p = p.Next {
		s = append(s, strconv.Itoa(p.Val))
	}
	for i := range s {
		curIndex := len(s) - i - 1
		output.WriteString(s[curIndex])
	}

	return output.String()
}

func compareIsEQ(l *ListNode, r *ListNode) bool {
	lp, rp := l, r
	for ; lp != nil && rp != nil; lp, rp = lp.Next, rp.Next {
		if lp.Val != rp.Val {
			return false
		}
	}
	if lp != nil || rp != nil {
		return false
	}
	return true
}

func TestAddTwoSum(t *testing.T) {
	testCases := []TestCase{
		TestCase{
			op1: generateListNode([]int{1, 2, 3}),
			op2: generateListNode([]int{4, 5, 6}),
			ans: generateListNode([]int{5, 7, 9}),
		},
		TestCase{
			op1: generateListNode([]int{4, 6, 3}),
			op2: generateListNode([]int{4, 5, 6}),
			ans: generateListNode([]int{8, 1, 0, 1}),
		},
		TestCase{
			op1: generateListNode([]int{8, 6, 3}),
			op2: generateListNode([]int{3, 5, 6, 5}),
			ans: generateListNode([]int{1, 2, 0, 6}),
		},
	}
	for _, curCase := range testCases {
		result := addTwoNumbers(curCase.op1, curCase.op2)
		isEqual := compareIsEQ(result, curCase.ans)
		if !isEqual {
			t.Errorf("Ans and result not match: calculated %v,"+
				"wanted %v",
				prettyPrintList(result),
				prettyPrintList(curCase.ans))

		}
	}

}
