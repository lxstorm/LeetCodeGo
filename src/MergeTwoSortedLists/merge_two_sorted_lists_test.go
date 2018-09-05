package MergeTwoSortedLists

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func pushBackNode(node *ListNode, val int) *ListNode {
	newNode := new(ListNode)
	newNode.Val = val
	newNode.Next = node.Next
	node.Next = newNode

	return newNode
}

func traversalNode(node *ListNode) {
	p := node
	var buffer []string
	for p != nil {
		buffer = append(buffer, strconv.Itoa(p.Val))
		p = p.Next
	}
	fmt.Println(strings.Join(buffer, "->"))
}

func TestF(t *testing.T) {
	dummy := new(ListNode)
	tail1 := dummy
	tail1 = pushBackNode(tail1, 1)
	tail1 = pushBackNode(tail1, 2)
	tail1 = pushBackNode(tail1, 4)
	l1 := dummy.Next
	dummy = new(ListNode)
	tail1 = dummy
	tail1 = pushBackNode(tail1, 1)
	tail1 = pushBackNode(tail1, 3)
	tail1 = pushBackNode(tail1, 4)
	l2 := dummy.Next

	traversalNode(l1)
	traversalNode(l2)
	result := mergeTwoLists(l1, l2)
	traversalNode(result)

}
