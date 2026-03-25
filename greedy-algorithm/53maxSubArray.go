package greedyalgorithm

/*
https://leetcode.cn/problems/maximum-subarray/description/

53. 最大子数组和

给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组是数组中的一个连续部分。
*/

func maxSubArray(nums []int) int {
	ans := nums[0]

	cur := 0

	for _, v := range nums {
		cur += v
		if cur > ans {
			ans = cur
		}
		if cur <= 0 {
			cur = 0
		}
	}

	return ans
}
