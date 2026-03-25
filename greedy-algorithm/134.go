package greedyalgorithm

/*
https://leetcode.cn/problems/gas-station/description/

134. 加油站

在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。

给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。
*/

func canCompleteCircuit(gas []int, cost []int) int {
	used := make([]int, len(gas))
	for i := 0; i < len(gas); i++ {
		used[i] = gas[i] - cost[i]
	}

	start := 0
	cur := 0
	totalSum := 0
	for i := 0; i < len(used); i++ {
		cur += used[i]
		totalSum += used[i]
		if cur < 0 {
			start = i + 1
			cur = 0
		}
	}

	if totalSum < 0 {
		return -1
	}

	return start
}
