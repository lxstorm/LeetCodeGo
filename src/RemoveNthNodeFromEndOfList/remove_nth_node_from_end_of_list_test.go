package RemoveNthNodeFromEndOfList

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
	head := new(ListNode)
	head.Val = 1
	head.Next = nil
	tail := head
	tail = pushBackNode(head, 2)
	tail = pushBackNode(tail, 3)
	tail = pushBackNode(tail, 4)
	tail = pushBackNode(tail, 5)

	traversalNode(head)
	result := removeNthFromEnd(head, 5)
	traversalNode(result)
}
