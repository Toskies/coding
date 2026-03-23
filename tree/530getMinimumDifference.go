package tree

import "math"

func getMinimumDifference(root *TreeNode) int {

	var prev *TreeNode
	curMin := math.MaxInt32
	var order func(root *TreeNode)
	order = func(root *TreeNode) {
		if root == nil {
			return
		}
		order(root.Left)
		if prev != nil && root.Val-prev.Val < curMin {
			curMin = root.Val - prev.Val
		}
		prev = root
		order(root.Right)
	}

	order(root)

	return curMin
}
