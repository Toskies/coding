package backtrack

/*
https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/

17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
*/

func letterCombinations(digits string) []string {
	m := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	ans := make([]string, 0)
	path := make([]byte, 0)

	var backtracking func(path []byte, start int)
	backtracking = func(path []byte, start int) {
		if len(path) == len(digits) {
			tmp := string(path)
			ans = append(ans, tmp)
			return
		}

		cur := int(digits[start] - '0')
		str := m[cur-2]
		for i := 0; i < len(str); i++ {
			path = append(path, str[i])
			backtracking(path, start+1)
			path = path[:len(path)-1]
		}
	}

	backtracking(path, 0)
	return ans
}
