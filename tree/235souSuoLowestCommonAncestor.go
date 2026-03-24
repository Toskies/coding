package tree

func souSuoLowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root.Val < p.Val && root.Val < q.Val {
		return souSuoLowestCommonAncestor(root.Right, p, q)
	}
	if root.Val > p.Val && root.Val > q.Val {
		return souSuoLowestCommonAncestor(root.Left, p, q)
	}

	return root
}
