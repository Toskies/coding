package greedyalgorithm

import "sort"

/*
https://leetcode.cn/problems/merge-intervals/description/

56. 合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
*/

func merge(intervals [][]int) [][]int {

	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := make([][]int, 0)
	start, end := intervals[0][0], intervals[0][1]

	for _, v := range intervals {
		if v[0] > end {
			ans = append(ans, []int{start, end})
			start = v[0]
			end = v[1]
		} else {
			start = min(start, v[0])
			end = max(end, v[1])
		}

	}
	ans = append(ans, []int{start, end})
	return ans
}
