package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
	n := len(cost)
	if n == 2 {
		return min(cost[0], cost[1])
	}
	pre, ppre := cost[1], cost[0]
	for i := 2; i < n; i++ {
		cur := cost[i] + min(pre, ppre)
		ppre = pre
		pre = cur
	}
	return pre
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	cost := []int{10, 15, 20}
	fmt.Printf("输入: %v\n", cost)
	r := minCostClimbingStairs(cost)
	fmt.Printf("输出: %v\n", r)
}
