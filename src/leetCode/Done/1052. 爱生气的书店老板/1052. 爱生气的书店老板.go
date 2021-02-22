package main

import "fmt"

func maxSatisfied(customers []int, grumpy []int, X int) int {
	n := len(customers)
	arr := make([]int, n)
	res := 0
	for i := range customers {
		if grumpy[i] == 1 {
			arr[i] = customers[i]
		} else {
			res += customers[i]
		}
	}
	preSum := make([]int, n+1)
	for i := range arr {
		preSum[i+1] = preSum[i] + arr[i]
	}
	if X >= n {
		return res + preSum[len(preSum)-1]
	}
	curValue := 0
	for R := X; R < len(preSum); R++ {
		curValue = max(curValue, preSum[R]-preSum[R-X])
	}
	return res + curValue
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
}
