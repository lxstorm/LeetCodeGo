package BinaryTreeInorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	_stack []*TreeNode
}

func (s Stack) empty() bool {
	if len(s._stack) == 0 {
		return true
	}
	return false
}

func (s *Stack) top() *TreeNode {
	return s._stack[len(s._stack)-1]
}

func (s *Stack) push(t *TreeNode) {
	s._stack = append(s._stack, t)
}

func (s *Stack) pop() {
	s._stack = s._stack[0 : len(s._stack)-1]
}

// stack use to remember the pre node
func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := new(Stack)
	p := root
	for p != nil || !stack.empty() {
		if p != nil {
			stack.push(p)
			p = p.Left
		} else {
			p = stack.top()
			result = append(result, p.Val)
			stack.pop()
			p = p.Right
		}
	}

	return result
}

// twice exercise
// func inorderTraversal(root *TreeNode) []int {
// 	result := []int{}
// 	stack := new(Stack)
// 	p := root
// 	for p != nil || !stack.empty() {
// 		if p != nil {
// 			stack.push(p)
// 			p = p.Left
// 		} else {
// 			p = stack.top()
// 			result = append(result, p.Val)
// 			stack.pop()
// 			p = p.Right
// 		}
// 	}
// 	return result
// }
