package backtrack

/*
https://leetcode.cn/problems/non-decreasing-subsequences/


491.递增子序列
给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。

示例:
输入: [4, 6, 7, 7]
输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]

输入：nums = [4,4,3,2,1]
输出：[[4,4]]

说明:
给定数组的长度不会超过15。
数组中的整数范围是 [-100,100]。
给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。
*/

func findSubsequences(nums []int) [][]int {
	/*
		错误方法：
		当前实现会重复返回结果，不符合题意。

		  根因在这里：  i > 0 && nums[i] == nums[i-1] && used[i-1] == false

		  这套去重条件是从“排序数组去重”那类题里借来的，只能处理“相同元素相邻”的情况。你的这个示例里，第一个 1 和后面那串 1 并不相邻，所以它们在同一层搜索里仍然会生成重复结果。

		  结论：当前文件逻辑有误，错误点是“去重策略不适用于本题”。本题应当在每一层递归里按“数值”去重，而不是按“前一个下标是否使用”去重。
	*/
	// ans := make([][]int, 0)
	// path := make([]int, 0)
	// used := make([]bool, len(nums))
	// var backtracking func(start int)
	// backtracking = func(start int) {
	// 	if len(path) >= 2 {
	// 		tmp := make([]int, len(path))
	// 		copy(tmp, path)
	// 		ans = append(ans, tmp)
	// 	}

	// 	if start == len(nums) {
	// 		return
	// 	}

	// 	for i := start; i < len(nums); i++ {
	// 		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
	// 			continue
	// 		}
	// 		if len(path) == 0 || nums[i] >= path[len(path)-1] {
	// 			path = append(path, nums[i])
	// 			used[i] = true
	// 			backtracking(i + 1)
	// 			used[i] = false
	// 			path = path[:len(path)-1]
	// 		}

	// 	}
	// }

	// backtracking(0)

	// return ans

	// 正确写法
	ans := make([][]int, 0)
	path := make([]int, 0)

	var backtracking func(start int)
	backtracking = func(start int) {
		if len(path) >= 2 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
		}

		used := make(map[int]bool, len(nums))
		for i := start; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}

			if len(path) == 0 || nums[i] >= path[len(path)-1] {
				path = append(path, nums[i])
				used[nums[i]] = true
				backtracking(i + 1)
				path = path[:len(path)-1]
			}

		}
	}

	backtracking(0)

	return ans
}
