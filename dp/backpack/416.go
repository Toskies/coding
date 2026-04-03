package backpack

/*
https://leetcode.cn/problems/partition-equal-subset-sum/description/

416. 分割等和子集
题目难易：中等

给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

1 <= nums.length <= 200
1 <= nums[i] <= 100
*/

func canPartition(nums []int) bool {
	sum := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false
	}
	target := sum / 2

	dp := make([]int, target+1)

	for i := 0; i <= target; i++ {
		if i >= nums[0] {
			dp[i] = nums[0]
		}
	}

	for i := 1; i < l; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}

	return dp[target] == target
}
