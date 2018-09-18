package RemoveDuplicatesFromSortedListII

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

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
	// nums := []int{1, 2, 3, 3, 4, 4, 5}
	nums := []int{1, 2, 2}
	ln := generateListNode(nums)
	traversalNode(ln)
	result := deleteDuplicates(ln)
	traversalNode(result)
}
