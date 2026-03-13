package tree

// 迭代
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		curLen := len(queue)
		temp := make([]int, 0)
		for curLen > 0 {
			index := queue[0]
			queue = queue[1:]

			temp = append(temp, index.Val)
			if index.Left != nil {
				queue = append(queue, index.Left)
			}
			if index.Right != nil {
				queue = append(queue, index.Right)
			}

			curLen--
		}
		ans = append(ans, temp)

	}

	return ans

}

// 递归

func levelOrderRecursion(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}

	depth := 0

	var oneLevel func(root *TreeNode, depth int)
	oneLevel = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if len(ans) == depth {
			ans = append(ans, []int{})
		}
		ans[depth] = append(ans[depth], root.Val)

		oneLevel(root.Left, depth+1)
		oneLevel(root.Right, depth+1)

	}

	oneLevel(root, depth)

	return ans
}
