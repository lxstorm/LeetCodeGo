package SwapNodesInPairs

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
	tail := dummy
	tail = pushBackNode(tail, 1)
	tail = pushBackNode(tail, 2)
	tail = pushBackNode(tail, 3)
	tail = pushBackNode(tail, 4)

	traversalNode(dummy.Next)
	result := swapPairs(dummy.Next)
	traversalNode(result)
}
