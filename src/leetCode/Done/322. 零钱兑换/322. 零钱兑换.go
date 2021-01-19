package main

import (
	"fmt"
	"math"
	"sort"
)

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	dp := make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		cur := math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if coins[j] > i {
				break
			}
			cur = min(cur, dp[i-coins[j]]+1)
		}
		dp[i] = cur
	}
	return dp[amount]
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x

}
func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}
