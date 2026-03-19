package tree

func sumOfLeftLeaves(root *TreeNode) int {
	ans := 0

	var order func(root *TreeNode, isLeft bool)
	order = func(root *TreeNode, isLeft bool) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil && isLeft {
			ans += root.Val
		}

		if root.Left != nil {
			order(root.Left, true)
		}
		if root.Right != nil {
			order(root.Right, false)
		}
	}

	order(root, false)

	return ans
}
