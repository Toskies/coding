package tree

func searchBST(root *TreeNode, val int) *TreeNode {
	cur := root
	for cur != nil {
		if cur.Val == val {
			return cur
		} else if cur.Val < val {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}

	return cur
}
