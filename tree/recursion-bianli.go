package tree

// 仅修改切片「已有元素的值」→ 不用传指针
// 因为副本的 ptr 指向同一个底层数组，修改元素值会同步到原切片：

// 修改切片「长度 / 容量」（如 append）→ 必须传指针
// append 时如果底层数组容量不足，会新建一个更大的数组，并让副本的 ptr 指向这个新数组 —— 但原切片的 sliceHeader 还是指向旧数组，导致修改「丢失」

func PreOrder(tree *TreeNode, ans *[]int) {
	if tree == nil {
		return
	}
	*ans = append(*ans, tree.Val)
	PreOrder(tree.Left, ans)
	PreOrder(tree.Right, ans)
}

func IneerOrder(tree *TreeNode, ans *[]int) {
	if tree == nil {
		return
	}

	IneerOrder(tree.Left, ans)
	*ans = append(*ans, tree.Val)
	IneerOrder(tree.Right, ans)
}

func PostOrder(tree *TreeNode, ans *[]int) {
	if tree == nil {
		return
	}

	PostOrder(tree.Left, ans)
	PostOrder(tree.Right, ans)
	*ans = append(*ans, tree.Val)
}

func Traversal(root *TreeNode) []int {
	ans := make([]int, 0)
	PreOrder(root, &ans)
	// IneerOrder(root, &ans)
	// PostOrder(root, &ans)
	return ans
}
