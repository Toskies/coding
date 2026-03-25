package backtrack

/*
https://leetcode.cn/problems/permutations-ii/description/

47. 全排列 II
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
*/

func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0, len(nums))
	used1 := make([]bool, len(nums))

	var backtracking func(cur int)
	backtracking = func(cur int) {
		if cur == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		used2 := make(map[int]bool, len(nums))
		for i := 0; i < len(nums); i++ {
			if used2[nums[i]] || used1[i] {
				continue
			}
			used2[nums[i]] = true
			path = append(path, nums[i])
			used1[i] = true
			backtracking(cur + 1)
			used1[i] = false
			path = path[:len(path)-1]

		}
	}

	backtracking(0)

	return ans
}
