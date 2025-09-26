/*
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果 target 存在返回下标，否则返回 -1。
你必须编写一个具有 O(log n) 时间复杂度的算法。

示例 1:

输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
示例 2:

输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1
*/

package array

func search(nums []int, target int) int {
	len := len(nums)
	index := (len - 1) / 2
	left := 0
	right := len - 1
	for left <= right {
		if nums[index] < target {
			left = index + 1
		} else if nums[index] > target {
			right = index - 1
		} else {
			return index
		}
		index = (left + right) / 2
	}
	return -1
}
