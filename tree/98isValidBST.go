package tree

func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	var travel func(node *TreeNode) bool
	travel = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		leftRes := travel(node.Left)
		if prev != nil && node.Val <= prev.Val {
			return false
		}
		prev = node
		rightRes := travel(node.Right)
		return leftRes && rightRes
	}
	return travel(root)
}
