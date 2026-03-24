package backtrack

/*
https://leetcode.cn/problems/palindrome-partitioning/description/

131. 分割回文串
给你一个字符串 s，请你将 s 分割成一些 子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
*/

func partition(s string) [][]string {
	ans := make([][]string, 0)
	path := make([]string, 0)

	var backtracking func(path []string, start int)
	backtracking = func(path []string, start int) {
		if start == len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		for i := start; i < len(s); i++ {
			str := s[start : i+1]
			if isPalindrome(str) {
				path = append(path, str)
				backtracking(path, i+1)
				path = path[:len(path)-1]
			}
		}
	}

	backtracking(path, 0)
	return ans
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
