/*
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。

示例：

输入：s = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
*/

package array

// 下面的解法是  求其和 等于s 的长度最小的 连续 子数组
// func minSubArrayLen(target int, nums []int) int {
// 	left, right := 0, 0
// 	sum, ans := 0, 0
// 	for right < len(nums) {
// 		sum += nums[right]
// 		if sum == target {
// 			if (right-left+1 < ans) || ans == 0 {
// 				ans = right - left + 1
// 			}
// 			sum -= nums[left]
// 			left++
// 			right++
// 		} else if sum > target {
// 			sum -= nums[left]
// 			sum -= nums[right]
// 			left++
// 		} else {
// 			right++
// 		}
// 	}

// 	return ans
// }

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum, ans := 0, 0
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			if (right-left+1 < ans) || ans == 0 {
				ans = right - left + 1
			}
			sum -= nums[left]
			left++
		}
		right++
	}

	return ans
}
