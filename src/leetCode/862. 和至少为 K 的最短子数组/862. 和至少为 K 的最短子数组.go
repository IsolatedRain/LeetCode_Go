package main

import "fmt"

func shortestSubarray(A []int, K int) int {
	preSum := []int{0}
	for i := range A {
		preSum = append(preSum, preSum[len(preSum)-1]+A[i])
	}
	fmt.Println(preSum)
	stack := []int{}
	res := len(preSum)
	for i := range preSum {
		for len(stack) > 0 && preSum[stack[len(stack)-1]] >= preSum[i] {
			stack = stack[:len(stack)-1]
		}
		for len(stack) > 0 && preSum[i]-preSum[stack[0]] >= K {
			fmt.Println(stack, res)
			res = min(res, i-stack[0])
			stack = stack[1:]
		}
		stack = append(stack, i)
	}
	return res
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x

}
func main() {
	fmt.Println(shortestSubarray([]int{17, 85, 93, -45, -21}, 150))
}
