package backtrack

/*
https://leetcode.cn/problems/combination-sum-iii/description/
216. 组合总和 III
找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：

只使用数字1到9
每个数字 最多使用一次
返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
*/

func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)

	var backtracking func(path []int, start, curSum int)
	backtracking = func(path []int, start, curSum int) {
		if len(path) == k && curSum == n {
			tmp := make([]int, k)
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		for i := start; i <= 9; i++ {
			if 10-i < k-len(path) || n-curSum < start {
				break
			}

			path = append(path, i)
			curSum += i
			backtracking(path, i+1, curSum)
			curSum -= i
			path = path[:len(path)-1]

		}
	}

	backtracking(path, 1, 0)

	return ans
}
