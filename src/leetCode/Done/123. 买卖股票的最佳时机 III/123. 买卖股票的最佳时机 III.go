package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	prices = append(prices, -1)
	n := len(prices)
	var forwardProfit func() []int
	forwardProfit = func() []int {
		r := make([]int, n)
		curMin := prices[0]
		for i := 1; i < n; i++ {
			r[i] = max(r[i-1], prices[i]-curMin)
			if curMin > prices[i] {
				curMin = prices[i]
			}
		}
		return r
	}
	forward := forwardProfit()
	var reversedProfit func() []int
	reversedProfit = func() []int {
		r := make([]int, n)
		curMax := prices[n-2]
		for i := n - 2; i > -1; i-- {
			r[i] = max(r[i+1], curMax-prices[i])
			curMax = max(prices[i], curMax)
		}
		return r
	}
	reversed := reversedProfit()
	res := 0
	for i := 1; i < n-1; i++ {
		res = max(res, forward[i]+reversed[i+1])
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
	// prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	prices := []int{2, 1, 4, 5, 2, 9, 7}
	fmt.Println(maxProfit(prices))
}
