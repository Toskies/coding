package backtrack

import "sort"

/*
https://leetcode.cn/problems/combination-sum/description/

39. 组合总和
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
对于给定的输入，保证和为 target 的不同组合数少于 150 个。
*/

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	ans := make([][]int, 0)
	path := make([]int, 0)

	var backtracking func(path []int, curSum, start int)
	backtracking = func(path []int, curSum, start int) {
		if curSum == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > curSum {
				break
			}

			path = append(path, candidates[i])
			backtracking(path, curSum-candidates[i], i)
			path = path[:len(path)-1]
		}

	}

	backtracking(path, target, 0)
	return ans
}
