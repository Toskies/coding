package greedyalgorithm

import (
	"strconv"
)

/*
https://leetcode.cn/problems/monotone-increasing-digits/description/

738. 单调递增的数字
当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。
给定一个整数 n ，返回 小于或等于 n 的最大数字，且数字呈 单调递增 。
*/

func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)

	for i := len(s) - 2; i >= 0; i-- {
		if s[i] > s[i+1] {
			s = s[:i] + string(s[i]-1) + s[i+1:]
			for j := i + 1; j < len(s); j++ {
				s = s[:j] + "9" + s[j+1:]
			}
		}
	}
	ans, _ := strconv.Atoi(s)
	return ans
}
