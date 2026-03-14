package tree

import "math"

func largestValues(root *TreeNode) []int {
	ans := make([]int, 0)

	if root == nil {
		return ans
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		l := len(stack)

		temp := math.MinInt
		for i := 0; i < l; i++ {
			if stack[i].Val > temp {
				temp = stack[i].Val
			}
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		ans = append(ans, temp)
		stack = stack[l:]
	}

	return ans
}
