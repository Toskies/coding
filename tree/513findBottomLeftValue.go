package tree

func findBottomLeftValue(root *TreeNode) int {
	ans := 0
	ansHeight := 0

	var order func(root *TreeNode, height int)
	order = func(root *TreeNode, height int) {
		if root == nil {
			return
		}

		if ansHeight < height {
			ans = root.Val
			ansHeight = height
		}

		if root.Left != nil {
			order(root.Left, height+1)
		}

		if root.Right != nil {
			order(root.Right, height+1)
		}
	}
	order(root, 1)
	return ans
}
