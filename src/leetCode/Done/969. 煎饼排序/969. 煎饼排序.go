// https://leetcode-cn.com/problems/pancake-sorting/
package main

import (
	"fmt"
	"sort"
)

func pancakeSort(arr []int) []int {
	nums := []int{}
	nums = append(nums, arr...)
	sort.Ints(nums)
	res := []int{}
	curI := len(nums) - 1
	for ; curI > 0; curI-- {
		for i := 0; i < curI; i++ {
			if i == curI {
				break
			}
			if nums[curI] == arr[i] {
				res = append(res, i+1)
				flip(arr, 0, i)
				res = append(res, curI+1)
				flip(arr, 0, curI)
			}
		}
	}
	return res
}

func flip(arr []int, L, R int) {
	for L < R {
		arr[L], arr[R] = arr[R], arr[L]
		L++
		R--
	}
}

func main() {
	arr := []int{3, 2, 4, 1}
	/*
		3 2 4 1
		4 2 3 1
		1 3 2 4
		3 1 2 4
		2 1 3 4
		1 2 3 4
	*/
	fmt.Println(pancakeSort(arr))
}
