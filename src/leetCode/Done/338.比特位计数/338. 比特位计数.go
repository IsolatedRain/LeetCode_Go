package main

import "fmt"

func countBits(num int) []int {
	dp := make([]int, num+1, num+1)
	fmt.Println(dp)
	for i := 1; i < num+1; i++ {
		dp[i] = dp[i-1&i] + 1
	}
	return dp
}

func main() {
	num := 5
	r := countBits(num)
	fmt.Println(r)
}
