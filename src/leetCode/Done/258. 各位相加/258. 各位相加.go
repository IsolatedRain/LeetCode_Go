// https://leetcode-cn.com/problems/add-digits/
package main

import "fmt"

func addDigits(num int) int {
	for num > 9 {
		x := 0
		for num > 0 {
			x += num % 10
			num /= 10
		}
		num = x
	}
	return num
}

func main() {
	fmt.Println(addDigits(38))
}
