package tree

import "math"

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	index := getMaxNode(nums)
	root := &TreeNode{Val: nums[index]}

	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])

	return root
}

func getMaxNode(nums []int) int {
	curMax := math.MinInt32
	index := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > curMax {
			index = i
			curMax = nums[i]
		}
	}

	return index
}
