package greedyalgorithm

import "sort"

/*
https://leetcode.cn/problems/non-overlapping-intervals/description/

435. 无重叠区间
给定一个区间的集合 intervals ，其中 intervals[i] = [starti, endi] 。返回 需要移除区间的最小数量，使剩余区间互不重叠 。

注意 只在一点上接触的区间是 不重叠的。例如 [1, 2] 和 [2, 3] 是不重叠的。
*/

func eraseOverlapIntervals(intervals [][]int) int {
	ans := 0
	cur := 0
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[cur][1] > intervals[i][0] {
			ans++
		} else {
			cur = i
		}
	}

	return ans
}
