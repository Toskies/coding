package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		curLen := len(queue)
		for i := 0; i < curLen; i++ {
			index := queue[i]
			if i < curLen-1 {
				index.Next = queue[i+1]
			} else {
				index.Next = nil
			}

			if index.Left != nil {
				queue = append(queue, index.Left)
			}
			if index.Right != nil {
				queue = append(queue, index.Right)
			}
		}
		queue = queue[curLen:]
	}

	return root
}
