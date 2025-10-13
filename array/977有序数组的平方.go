/*
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

示例 1：

输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]
*/
package array

func sortedSquares(nums []int) []int {
	ans := []int{}
	left, right := 0, len(nums)-1
	for left <= right {
		l, r := nums[left]*nums[left], nums[right]*nums[right]
		if l < r {
			ans = append(ans, r)
			right--
		} else {
			ans = append(ans, l)
			left++
		}
	}

	left, right = 0, len(ans)-1
	for left < right {
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}
	return ans
}
