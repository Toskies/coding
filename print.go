package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// bufio中读取数据的接口，因为数据卡的比较严，导致使用fmt.Scan会超时
	scanner := bufio.NewScanner(os.Stdin)

	// 获取数组大小
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// 获取数组元素的同时计算前缀和，一般建议切片开大一点防止各种越界问题
	arr := make([]int, n+1)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
		if i != 0 {
			arr[i] += arr[i-1]
		}
	}

	/*
	   区间[l, r]的和可以使用区间[0, r]和[0, l - 1]相减得到，
	   在代码中即为arr[r]-arr[l-1]。这里需要注意l-1是否越界
	*/
	for {
		var l, r int
		scanner.Scan()
		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &l, &r)
		if err != nil {
			return
		}

		if l > 0 {
			fmt.Println(arr[r] - arr[l-1])
		} else {
			fmt.Println(arr[r])
		}
	}
}
