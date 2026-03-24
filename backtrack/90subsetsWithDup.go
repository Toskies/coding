package backtrack

import "sort"

/*
https://leetcode.cn/problems/subsets-ii/description/

90. 子集 II

给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的 子集（幂集）。
解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
*/
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))

	var backtracking func(start int)
	backtracking = func(start int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		ans = append(ans, tmp)

		for i := start; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			backtracking(i + 1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	backtracking(0)
	return ans
}
