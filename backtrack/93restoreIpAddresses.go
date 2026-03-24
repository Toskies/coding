package backtrack

import (
	"strconv"
	"strings"
)

/*
https://leetcode.cn/problems/restore-ip-addresses/description/

93. 复原 IP 地址

给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效的 IP 地址。

示例 1：
输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]
*/

func restoreIpAddresses(s string) []string {
	ans := make([]string, 0)
	path := make([]string, 0)

	var backtracking func(start int)
	backtracking = func(start int) {
		if len(path) == 4 { // 够四段后就不再继续往下递归
			if start == len(s) {
				str := strings.Join(path, ".")
				ans = append(ans, str)
			}
			return
		}

		for i := start; i < len(s); i++ {
			if s[start] == '0' { // 含有前导 0，无效
				break
			}
			str := s[start : i+1]
			num, _ := strconv.Atoi(str)
			if num >= 0 && num <= 255 {
				path = append(path, str) // 符合条件的就进入下一层
				backtracking(i + 1)
				path = path[:len(path)-1]
			} else { // 如果不满足条件，再往后也不可能满足条件，直接退出
				break
			}
		}
	}

	backtracking(0)

	return ans
}
