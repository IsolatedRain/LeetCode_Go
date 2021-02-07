package main

import (
	"fmt"
	"sort"
)

func maxValue(events [][]int, k int) int {
	n := len(events)
	sort.Slice(events, func(i, j int) bool { return events[i][1] < events[j][1] })
	R := make([]int, n)
	for i := range R {
		R[i] = events[i][1]
	}
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][1] = events[0][2]
	for i := 2; i < n+1; i++ {
		dp[0][i] = max(dp[0][i-1], events[i-1][2])
	}
	for i := 1; i < k; i++ {
		dp[i][1] = dp[i-1][1]
	}
	for r := 1; r < k; r++ {
		for c := 2; c < n+1; c++ {
			id := sort.Search(n, func(i int) bool { return R[i] >= events[c-1][0] })
			dp[r][c] = max(dp[r][c-1], dp[r-1][id]+events[c-1][2])
		}
	}
	return dp[k-1][n]
}
func max(a ...int) int {
	r := a[0]
	for _, v := range a[1:] {
		if v > r {
			r = v
		}
	}
	return r
}

func main() {
	events := [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}
	k := 2
	fmt.Println(maxValue(events, k))

}
