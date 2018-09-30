package UniqueBinarySearchTreesII

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	var helper func(begin, end int) []*TreeNode
	helper = func(begin, end int) []*TreeNode {
		res := []*TreeNode{}
		if begin > end {
			res = append(res, nil)
			return res
		}
		for i := begin; i <= end; i++ {
			// left tree from begin to i-1
			leftTreeList := helper(begin, i-1)
			rightTreeList := helper(i+1, end)
			for j := 0; j < len(leftTreeList); j++ {
				for k := 0; k < len(rightTreeList); k++ {
					root := new(TreeNode)
					root.Val = i
					root.Left = leftTreeList[j]
					root.Right = rightTreeList[k]
					res = append(res, root)
				}
			}
		}
		return res
	}

	return helper(1, n)
}

// another dp solution, change in the n-1 tree
// 1) n is the root
// 2) n is not the root, if the node has a right child, old node.Right = nth node
// 	nth node.Left = old node.Left
