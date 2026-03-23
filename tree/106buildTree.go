package tree

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	curRootValue := postorder[len(postorder)-1]
	root := &TreeNode{Val: curRootValue}

	if len(postorder) == 1 {
		return root
	}

	var midIndex int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == curRootValue {
			midIndex = i
			break
		}
	}

	root.Left = buildTree(inorder[:midIndex], postorder[:midIndex])
	root.Right = buildTree(inorder[midIndex+1:], postorder[midIndex:len(postorder)-1])

	return root
}
