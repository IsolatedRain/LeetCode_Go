// https://leetcode-cn.com/problems/next-greater-element-ii/
package main

import "fmt"

// 单调递减栈，维护到当前的 比当前大的元素。

func nextGreaterElements(nums []int) []int {
	arr := append([]int{}, nums...)
	arr = append(arr, nums...)
	n := len(arr)
	if n == 0 {
		return []int{}
	}
	res := make([]int, n)
	res[n-1] = -1
	stack := []int{arr[n-1]}
	for i := n - 2; i > -1; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}
		stack = append(stack, arr[i])
	}
	return res[:len(nums)]
}

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 5, 3, 4}))
}
