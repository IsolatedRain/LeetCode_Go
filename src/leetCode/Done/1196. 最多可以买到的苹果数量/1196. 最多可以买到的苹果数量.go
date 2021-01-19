package main

import (
	"fmt"
	"sort"
)

func maxNumberOfApples(arr []int) int {
	sort.Ints(arr)
	cnt := 0
	curWeight := 0
	for _, v := range arr {
		if curWeight += v; curWeight > 5000 {
			break
		}
		cnt++
	}
	return cnt
}

func main() {
	arr := []int{900, 950, 800, 1000, 700, 800}
	fmt.Println(maxNumberOfApples(arr))
}
