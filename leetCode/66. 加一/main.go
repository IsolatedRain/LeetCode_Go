package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i > -1; i-- {
		// 进位最多为 1, 如果小于 9, 当前位置加 1, 其他不变返回.
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	// 都为 9999..., 加一后 为 0000...
	// 修改第一位为1, 后面 追加 0
	digits[0] = 1
	return append(digits, 0)
}

func main() {
	// p := []int{4, 3, 2, 1}
	p := []int{9, 9, 9, 9}
	res := plusOne(p)
	fmt.Printf("%v\n", res)
}
