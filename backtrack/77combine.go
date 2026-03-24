package backtrack

/*
https://leetcode.cn/problems/combinations/description/

77. 组合
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。
*/

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)

	var backtracking func(path []int, start int)

	backtracking = func(path []int, start int) {
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		for i := start; i <= n; i++ {
			if n-i+1 < k-len(path) {
				break
			}

			path = append(path, i)
			backtracking(path, i+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(path, 1)

	return ans
}
