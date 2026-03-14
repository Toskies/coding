package tree

func invertTree(root *TreeNode) *TreeNode {
	var order func(root *TreeNode)
	order = func(root *TreeNode) {
		if root == nil {
			return
		}

		root.Left, root.Right = root.Right, root.Left
		order(root.Left)
		order(root.Right)
	}

	order(root)
	return root
}
