package greedyalgorithm

import (
	"coding/tree"
	"math"
)

/*
https://leetcode.cn/problems/binary-tree-cameras/description/

968. 监控二叉树
给定一个二叉树，我们在树的节点上安装摄像头。

节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。

计算监控树的所有节点所需的最小摄像头数量。
*/

/*
a：当前装摄像头，子树全覆盖
b：当前不装但被覆盖，子树全覆盖
c：当前不装且未被覆盖，子树全覆盖

 1. a = 当前节点装摄像头
    装了摄像头，自己和左右孩子一定都被照到
    左右孩子随便什么状态都可以，取最小就行
    最后 +1（加上当前这个摄像头）
    a = min(左a,左b,左c) + min(右a,右b,右c) + 1
 2. b = 当前节点不装，但被覆盖
    自己没摄像头，必须靠孩子照亮自己
    所以：左孩子装 或 右孩子装，至少一个
    两种合法情况取最小：
    左装 + 右被覆盖
    右装 + 左被覆盖
    b = min(左a+右b, 右a+左b)
 3. c = 当前节点不装，也没被覆盖
    自己没被照到，但孩子必须都被覆盖好
    孩子不能是 “没被覆盖” 的状态，否则整棵树不合法
    所以只能：左被覆盖 + 右被覆盖
    c = 左b + 右b

一句话再浓缩

	装摄像头：孩子随便，+1
	不装但被照：至少一个孩子装
	不装也没被照：两个孩子都必须被照好
*/
const inf = math.MaxInt64 / 2

func minCameraCover(root *tree.TreeNode) int {
	var dfs func(*tree.TreeNode) (a, b, c int)
	dfs = func(node *tree.TreeNode) (a, b, c int) {
		if node == nil {
			return inf, 0, 0
		}
		lefta, leftb, leftc := dfs(node.Left)
		righta, rightb, rightc := dfs(node.Right)
		a = leftc + rightc + 1
		b = min(a, min(lefta+rightb, righta+leftb))
		c = min(a, leftb+rightb)
		return
	}
	_, ans, _ := dfs(root)
	return ans
}
