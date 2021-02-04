package main

import (
	"fmt"
)

func equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
	diff := []int{}
	for i := range s {
		diff = append(diff, abs(s[i], t[i]))
	}
	res, L, R, curCost := 0, 0, 0, 0
	for R < n {
		curCost += diff[R]
		for curCost > maxCost {
			curCost -= diff[L]
			L++
		}
		R++
		res = max(res, R-L)
	}
	return res
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func abs(a, b byte) int {
	if a > b {
		return int(a - b)
	}
	return int(b - a)
}

func main() {
	// s := "abcd"
	// t := "bcdf"
	// cost := 3
	s := "tuhwpmkerswypuexja"
	t := "sdmxfcgynmhhfvwcfp"
	cost := 173
	fmt.Println(equalSubstring(s, t, cost))
}

// 二分
// func equalSubstring(s string, t string, maxCost int) int {
// 	n := len(s)
// 	diff := []int{}
// 	for i := range s {
// 		diff = append(diff, abs(s[i], t[i]))
// 	}
// 	preSum := []int{0}
// 	for i := range diff {
// 		preSum = append(preSum, preSum[len(preSum)-1]+diff[i])
// 	}
// 	if preSum[len(preSum)-1] <= maxCost {
// 		return n
// 	}
// 	res := 0
// 	for i := range preSum {
// 		if preSum[i] >= maxCost {
// 			L := sort.Search(n, func(j int) bool { return preSum[j] >= preSum[i]-maxCost })
// 			res = max(res, i-L)
// 		} else {
// 			res = max(res, i)
// 		}
// 	}
// 	return res
// }
