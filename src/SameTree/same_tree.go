package SameTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursion method
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// dfs pre-order method
func isSameTreeDFS(p *TreeNode, q *TreeNode) bool {
	v := []*TreeNode{}
	v = append(v, p)
	v = append(v, q)
	for len(v) != 0 {
		pTop := v[len(v)-2]
		qTop := v[len(v)-1]
		v = v[:len(v)-2]
		if pTop == nil && qTop == nil {
			continue
		}
		if pTop == nil || qTop == nil || pTop.Val != qTop.Val {
			return false
		}
		v = append(v, pTop.Left)
		v = append(v, qTop.Left)
		v = append(v, pTop.Right)
		v = append(v, qTop.Right)
	}
	return true

}
