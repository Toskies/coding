package tree

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := make([]int, 0)

	stack := make([]*TreeNode, 0)
	index := root
	stack = append(stack, index)

	for len(stack) > 0 {
		index = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		ans = append(ans, index.Val)

		if index.Right != nil {
			stack = append(stack, index.Right)
		}
		if index.Left != nil {
			stack = append(stack, index.Left)
		}
	}

	return ans
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := make([]int, 0)

	stack := make([]*TreeNode, 0)
	index := root
	stack = append(stack, index)

	for len(stack) > 0 || index != nil {
		if index != nil {
			stack = append(stack, index)
			index = index.Left
		} else {
			index = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans = append(ans, index.Val)
			index = index.Right
		}

	}

	return ans
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		index := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		ans = append(ans, index.Val)
		if index.Left != nil {
			stack = append(stack, index.Left)
		}
		if index.Right != nil {
			stack = append(stack, index.Right)
		}
	}

	resverse(ans)
	return ans
}

func resverse(ans []int) []int {
	left, right := 0, len(ans)-1

	for left < right {
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}

	return ans
}
