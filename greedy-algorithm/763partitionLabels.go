package greedyalgorithm

/*
https://leetcode.cn/problems/partition-labels/description/

763. 划分字母区间
给你一个字符串 s 。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。
例如，字符串 "ababcc" 能够被分为 ["abab", "cc"]，但类似 ["aba", "bcc"] 或 ["ab", "ab", "cc"] 的划分是非法的。

注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。

返回一个表示每个字符串片段的长度的列表。

示例：

输入：S = "ababcbacadefegdehijhklij"
输出：[9,7,8] 解释： 划分结果为 "ababcbaca", "defegde", "hijhklij"。 每个字母最多出现在一个片段中。 像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。

提示：

S的长度在[1, 500]之间。
S只包含小写字母 'a' 到 'z' 。
*/

func partitionLabels(s string) []int {
	if len(s) == 0 {
		return []int{}
	}

	index := make([]int, 26)
	for i, v := range s {
		if i > index[v-'a'] {
			index[v-'a'] = i
		}
	}

	ans := make([]int, 0)
	cur, pre := index[s[0]-'a'], 0

	for i, v := range s {
		cur = max(cur, index[v-'a'])

		if cur == i {
			ans = append(ans, cur-pre+1)
			pre = cur + 1
			continue
		}
		if cur == len(s)-1 {
			ans = append(ans, cur-pre+1)
			break
		}
	}

	return ans
}
