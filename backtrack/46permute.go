package backtrack

/*
https://leetcode.cn/problems/permutations/description/

46. 全排列
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:
输入: [1,2,3]
输出: [ [1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1] ]
*/

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0, len(nums))

	var backtracking func(cur []int)
	backtracking = func(cur []int) {
		if len(cur) == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		for i := 0; i < len(cur); i++ {
			path = append(path, cur[i])
			next := make([]int, 0, len(cur)-1)
			next = append(next, cur[:i]...)
			next = append(next, cur[i+1:]...)
			backtracking(next)
			path = path[:len(path)-1]
		}

	}

	backtracking(nums)

	return ans
}
