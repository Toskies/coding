package tree

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	ans := make([]string, 0)
	if root == nil {
		return ans
	}

	var order func(root *TreeNode, old string)
	order = func(root *TreeNode, old string) {
		if root.Left == nil && root.Right == nil {
			ans = append(ans, old)
			return
		}

		if root.Left != nil {
			order(root.Left, old+"->"+strconv.Itoa(root.Left.Val))
		}
		if root.Right != nil {
			order(root.Right, old+"->"+strconv.Itoa(root.Right.Val))
		}
	}

	order(root, strconv.Itoa(root.Val))

	return ans
}
