/*
给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3 输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]
*/

package array

func generateMatrix(n int) [][]int {
	loop, mid := n/2, n/2
	conut := 1
	offset := 1
	startx, starty := 0, 0
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	for loop > 0 {

		for i := starty; i < n-offset; i++ {
			ans[startx][i] = conut
			conut++
		}

		for j := startx; j < n-offset; j++ {
			ans[j][n-offset] = conut
			conut++
		}

		for i := n - offset; i > starty; i-- {
			ans[n-offset][i] = conut
			conut++
		}

		for j := n - offset; j > startx; j-- {
			ans[j][starty] = conut
			conut++
		}

		startx++
		starty++
		offset++
		loop--
	}

	if n%2 == 1 {
		ans[mid][mid] = conut
	}
	return ans
}
