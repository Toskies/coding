package backtrack

/*
https://leetcode.cn/problems/subsets/description/

78. 子集
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
*/

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)

	var backtracking func(start int)
	backtracking = func(start int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		ans = append(ans, tmp)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)
	return ans
}
