package main

import (
	"fmt"
	"sort"
)

func lengthOfLIS(nums []int) int {
	stack := []int{nums[0]}
	for _, num := range nums[1:] {
		if stack[len(stack)-1] < num {
			stack = append(stack, num)
		} else {
			idx := sort.Search(len(stack), func(i int) bool {
				return stack[i] >= num
			})
			stack[idx] = num
		}
	}
	return len(stack)
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}
