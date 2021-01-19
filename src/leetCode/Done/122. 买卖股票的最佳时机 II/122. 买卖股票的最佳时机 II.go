package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	stack := []int{}
	profit := 0
	for i := range prices {
		if len(stack) == 0 {
			stack = append(stack, prices[i])
		} else {
			if prices[i] < stack[len(stack)-1] {
				profit += stack[len(stack)-1] - stack[0]
				stack = []int{prices[i]}
			} else {
				stack = append(stack, prices[i])
			}
		}
	}
	if len(stack) > 1 {
		profit += stack[len(stack)-1] - stack[0]
	}
	return profit
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
