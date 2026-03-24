package tree

func findMode(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	curMaxNum := 1
	curNum := 1
	// prev := root
	var prev *TreeNode

	var traval func(root *TreeNode)
	traval = func(root *TreeNode) {
		if root == nil {
			return
		}

		traval(root.Left)
		if prev != nil && prev.Val == root.Val {
			curNum++
		} else {
			curNum = 1
		}

		if curNum >= curMaxNum {
			if curNum > curMaxNum && len(ans) > 0 {
				ans = []int{root.Val}
			} else {
				ans = append(ans, root.Val)
			}
			curMaxNum = curNum
		}

		prev = root
		traval(root.Right)
	}

	traval(root)

	return ans
}
