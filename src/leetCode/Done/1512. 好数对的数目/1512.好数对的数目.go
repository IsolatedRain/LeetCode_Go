package main

import (
	"fmt"
)

func numIdenticalPairs(nums []int) int {
	count := make([]int, 101, 101)
	res := 0
	for _, i := range nums {
		res += count[i]
		count[i]++
	}
	// count := map[int]int{}
	// for _, v := range nums {
	// 	count[v]++
	// }
	// fmt.Printf("%v\n", count)
	// res := 0
	// for _, v := range count {
	// 	if v > 1 {
	// 		res += v * (v - 1) / 2
	// 	}
	// }

	return res
}

func main() {
	p := []int{1, 2, 3, 1, 1, 3}
	r := numIdenticalPairs(p)
	fmt.Printf("输入: %v\n输出: %v\n", p, r)
}
