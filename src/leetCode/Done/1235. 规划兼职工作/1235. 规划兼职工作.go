package main

import (
	"fmt"
	"sort"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	jobs := make([][]int, n)
	for i := range jobs {
		jobs[i] = []int{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool { return jobs[i][1] < jobs[j][1] })
	dp := [][]int{{0, 0}}
	for i := range jobs {
		id := sort.Search(len(dp), func(j int) bool { return dp[j][0] > jobs[i][0] }) - 1
		if dp[id][1]+jobs[i][2] > dp[len(dp)-1][1] {
			dp = append(dp, []int{jobs[i][1], jobs[i][2] + dp[id][1]})
		}
	}
	return dp[len(dp)-1][1]
}

func main() {
	startTime := []int{1, 2, 3, 3}
	endTime := []int{3, 4, 5, 6}
	profit := []int{50, 10, 40, 70}
	fmt.Println(jobScheduling(startTime, endTime, profit))
}
