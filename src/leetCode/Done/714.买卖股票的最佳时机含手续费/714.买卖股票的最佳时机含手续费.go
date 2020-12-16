package main

import "fmt"

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	hold := -prices[0]
	sold := 0
	for i := 1; i < n; i++ {
		hold = max(hold, sold-prices[i])
		sold = max(sold, hold+prices[i]-fee)
	}
	return sold
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	r := maxProfit(prices, fee)
	fmt.Println(r)
}
