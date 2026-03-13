package tree

func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		curLen := len(queue)
		for curLen > 0 {
			index := queue[0]
			queue = queue[1:]

			if index.Left != nil {
				queue = append(queue, index.Left)
			}
			if index.Right != nil {
				queue = append(queue, index.Right)
			}

			curLen--
			if curLen == 0 {
				ans = append(ans, index.Val)
			}
		}

	}

	return ans
}
