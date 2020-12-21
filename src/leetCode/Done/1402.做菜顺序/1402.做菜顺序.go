package main

import (
	"fmt"
	"sort"
)

// 倒序，取最大值， 如果小于0，即不能增加了，就返回。
func maxSatisfaction(satisfaction []int) int {
	sort.Slice(satisfaction, func(i, j int) bool {
		return satisfaction[i] > satisfaction[j]
	})
	res := 0
	curSum := 0
	for _, s := range satisfaction {
		if curSum+s >= 0 {
			curSum += s
		} else {
			break
		}
		res += curSum
	}
	return res
}

func main() {
	satisfaction := []int{-1, -8, 0, 5, -9}
	r := maxSatisfaction(satisfaction)
	fmt.Println(r)
}
