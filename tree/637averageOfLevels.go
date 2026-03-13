package tree

func averageOfLevels(root *TreeNode) []float64 {
	ans := make([]float64, 0)
	if root == nil {
		return ans
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		curLen, tempLen := len(queue), len(queue)
		curNum := 0
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
			curNum += index.Val
			if curLen == 0 {
				ans = append(ans, float64(curNum)/float64(tempLen))
			}
		}

	}

	return ans
}
