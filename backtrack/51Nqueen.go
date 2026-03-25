package backtrack

/*

https://leetcode.cn/problems/n-queens/description/

51. N 皇后

按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
*/

func solveNQueens(n int) [][]string {
	ans := make([][]string, 0)
	path := make([]string, 0)
	usedCol := make([]bool, n)
	usedDiag1 := make([]bool, 2*n-1)
	usedDiag2 := make([]bool, 2*n-1)

	var backtracking func(cur int)
	backtracking = func(cur int) {
		if cur == n {
			tmp := make([]string, n)
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		for col := 0; col < n; col++ {
			diag1 := cur - col + n - 1
			diag2 := cur + col

			if usedCol[col] || usedDiag1[diag1] || usedDiag2[diag2] {
				continue
			}

			path = append(path, constructQueen(n, col))
			usedCol[col] = true
			usedDiag1[diag1] = true
			usedDiag2[diag2] = true
			backtracking(cur + 1)
			usedCol[col] = false
			usedDiag1[diag1] = false
			usedDiag2[diag2] = false
			path = path[:len(path)-1]
		}

	}

	backtracking(0)
	return ans
}

func constructQueen(n, index int) string {
	ans := []byte{}
	for i := 0; i < n; i++ {
		if i != index {
			ans = append(ans, '.')
		} else {
			ans = append(ans, 'Q')
		}
	}

	return string(ans)
}
