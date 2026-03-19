package tree

func hasPathSum(root *TreeNode, targetSum int) bool {
	var order func(root *TreeNode, sum int) bool
	order = func(root *TreeNode, sum int) bool {
		if root == nil {
			return false
		}
		if root.Left == nil && root.Right == nil {
			if sum+root.Val == targetSum {
				return true
			}
		}

		return order(root.Left, sum+root.Val) || order(root.Right, sum+root.Val)
	}

	return order(root, 0)
}
