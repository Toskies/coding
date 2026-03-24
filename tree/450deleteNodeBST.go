package tree

/*
https://leetcode.cn/problems/delete-node-in-a-bst/description/

450.删除二叉搜索树中的节点

给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。
*/

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	if root.Right == nil {
		return root.Left
	}
	if root.Left == nil {
		return root.Right
	}

	minRight := root.Right
	for minRight.Left != nil {
		minRight = minRight.Left
	}
	root.Val = minRight.Val
	root.Right = deleteNode1(root.Right)
	return root
}

func deleteNode1(root *TreeNode) *TreeNode {
	if root.Left == nil {
		tmp := root.Right
		// root.Right = nil   都可以
		return tmp
	}
	root.Left = deleteNode1(root.Left)
	return root
}
