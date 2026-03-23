package tree

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	var merge func(r1, r2 *TreeNode) *TreeNode
	merge = func(r1, r2 *TreeNode) *TreeNode {
		if r1 == nil {
			return r2
		}

		if r2 == nil {
			return r1
		}

		r1.Val += r2.Val
		r1.Left = merge(r1.Left, r2.Left)
		r1.Right = merge(r1.Right, r2.Right)

		return r1
	}

	return merge(root1, root2)
}
