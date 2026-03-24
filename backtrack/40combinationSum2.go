package backtrack

import "sort"

/*
https://leetcode.cn/problems/combination-sum-ii/description/

40. 组合总和 II
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用 一次 。
注意：解集不能包含重复的组合。
*/

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	ans := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(candidates))
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

			if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
				continue
			}

			path = append(path, candidates[i])
			used[i] = true
			backtracking(path, curSum-candidates[i], i+1)
			used[i] = false
			path = path[:len(path)-1]
		}

	}

	backtracking(path, target, 0)
	return ans
}
