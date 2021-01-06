package main

import (
	"fmt"
	"sort"
)

func minOperations(target []int, arr []int) int {
	idxMap := map[int]int{}
	for i, v := range target {
		idxMap[v] = i
	}
	nums := []int{}
	for _, v := range arr {
		if i, ok := idxMap[v]; ok {
			nums = append(nums, i)
		}
	}
	stack := []int{}
	for _, v := range nums {
		if len(stack) == 0 || stack[len(stack)-1] < v {
			stack = append(stack, v)
		} else {
			i := sort.Search(len(stack), func(j int) bool {
				return stack[j] >= v
			})
			stack[i] = v
		}
	}
	fmt.Println(nums)
	fmt.Println(stack)
	return len(target) - len(stack)
}

func main() {
	// target := []int{6, 4, 8, 1, 3, 2}
	// arr := []int{4, 7, 6, 2, 3, 8, 6, 1}
	target := []int{17, 18, 14, 13, 6, 9, 1, 3, 2, 20}
	arr := []int{18, 15, 14, 6, 6, 13, 15, 20, 2, 6}
	r := minOperations(target, arr)
	fmt.Println(r)
}
