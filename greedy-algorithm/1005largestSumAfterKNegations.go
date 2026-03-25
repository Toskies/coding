package greedyalgorithm

import (
	"math"
	"sort"
)

/*
https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations/description/

1005. K 次取反后最大化的数组和
给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：

选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
重复这个过程恰好 k 次。可以多次选择同一个下标 i 。

以这种方式修改数组后，返回数组 可能的最大和 。
*/

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)

	sum := 0
	curMin := math.MaxInt

	for i := 0; i < len(nums); i++ {
		if k > 0 && nums[i] < 0 {
			k--
			nums[i] = -nums[i]
		}
		sum += nums[i]
		curMin = min(curMin, nums[i])
	}

	if k != 0 {
		if k%2 == 1 {
			sum -= curMin * 2
		}
	}
	return sum
}
