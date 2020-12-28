package main

import (
	"fmt"
)

func maxProfit(k int, prices []int) int {
	n := len(prices)
	k = min(k, n/2)
	minNum := -10000000
	buyIn := make([]int, k+1)
	sellOut := make([]int, k+1)
	buyIn[0] = -prices[0]
	for i := 1; i < k+1; i++ {
		buyIn[i], sellOut[i] = minNum, minNum
	}
	for i := 1; i < n; i++ {
		buyIn[0] = max(buyIn[0], sellOut[0]-prices[i])
		for j := 1; j <= k; j++ {
			buyIn[j] = max(buyIn[j], sellOut[j]-prices[i])
			sellOut[j] = max(sellOut[j], buyIn[j-1]+prices[i])
		}
	}
	return max(sellOut...)
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func max(x ...int) int {
	r := x[0]
	for _, v := range x[1:] {
		if r < v {
			r = v
		}
	}
	return r
}

func main() {
	// k := 2
	// prices := []int{3, 2, 6, 5, 0, 3}
	// k := 2
	// prices := []int{2, 4, 1}
	k := 2
	prices := []int{3, 2, 6, 5, 0, 3, 213, 4, 5, 6, 7, 123, 4, 5, 6, 12, 3, 4, 5, 6, 7, 123, 4, 5, 7}
	r := maxProfit(k, prices)
	fmt.Println(r)
}

// 二维DP
// func maxProfit(k int, prices []int) int {
// 	if k == 0 {
// 		return 0
// 	}
// 	n := len(prices)
// 	minNum := -1000000
// 	hold := make([][]int, n)
// 	sold := make([][]int, n)
// 	for i := 0; i < n; i++ {
// 		hold[i] = make([]int, k)
// 		sold[i] = make([]int, k)
// 		for j := 0; j < k; j++ {
// 			hold[i][j] = minNum
// 			sold[i][j] = minNum
// 		}
// 	}
// 	hold[0][0] = -prices[0]
// 	for j := 0; j < k; j++ {
// 		sold[0][j] = minNum
// 	}
// 	for i := 1; i < n; i++ {
// 		hold[i][0] = max(-prices[i], hold[i-1][0])
// 		sold[i][0] = max(prices[i]+hold[i-1][0], sold[i-1][0])
// 	}
// 	for i := 1; i < n; i++ {
// 		for j := 1; j < k; j++ {
// 			hold[i][j] = max(sold[i-1][j-1]-prices[i], hold[i-1][j])
// 			sold[i][j] = max(hold[i-1][j]+prices[i], sold[i-1][j])
// 		}
// 	}

// 	return max(max(sold[n-1]...), 0)
// }
