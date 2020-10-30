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
	sufMax := prices[n-1] // 从后 到 当前位置的 最大值
	res := 0
	for i := n - 2; i > -1; i-- {
		sufMax = max(prices[i], sufMax)  // 更新最大值
		res = max(res, sufMax-prices[i]) // 更新最大差值
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
	p := []int{7, 1, 5, 3, 6, 4}
	r := maxProfit(p)
	fmt.Printf("输入: %v\n输出: %v\n", p, r)
}
