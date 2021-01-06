package main

import "fmt"

// 从后往前保留最大,
// 从前往后, 保留最小.
//去差值
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	res := 0
	preMin := prices[0]
	for i := 1; i < n; i++ {
		res = max(res, prices[i]-preMin)
		if preMin > prices[i] {
			preMin = prices[i]
		}
	}
	return res
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	p := []int{7, 1, 5, 3, 6, 9}
	r := maxProfit(p)
	fmt.Printf("输入: %v\n输出: %v\n", p, r)
}
