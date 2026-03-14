package tree

type NNode struct {
	Val      int
	Children []*NNode
}

// 迭代
func NlevelOrder(root *NNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}

	stack := make([]*NNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		l := len(stack)
		temp := make([]int, l)

		for i := 0; i < l; i++ {
			temp[i] = stack[i].Val

			for _, v := range stack[i].Children {
				stack = append(stack, v)
			}
		}
		stack = stack[l:]
		ans = append(ans, temp)
	}

	return ans
}

// 递归

func NDlevelOrder(root *NNode) [][]int {
	ans := make([][]int, 0)

	if root == nil {
		return ans
	}
	depth := 0

	var Norder func(root *NNode, depth int)
	Norder = func(root *NNode, depth int) {
		if root == nil {
			return
		}

		if len(ans) == depth {
			ans = append(ans, []int{})
		}
		ans[depth] = append(ans[depth], root.Val)
		for _, v := range root.Children {
			Norder(v, depth+1)
		}
	}

	Norder(root, depth)

	return ans

}
