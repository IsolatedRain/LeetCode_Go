package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, v := range tokens {
		if isDigit(v) {
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		} else {
			x, y := stack[len(stack)-2], stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], cal(x, y, v))
		}
	}
	return stack[0]
}
func cal(x, y int, op string) int {
	switch op {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		return x / y
	}
	return -1
}
func isDigit(s string) bool {
	if s[0] == '-' && len(s) > 1 {
		s = s[1:]
	}
	for i := range s {
		if '0' > s[i] || s[i] > '9' {
			return false
		}
	}
	return true
}

func main() {
	tokens := []string{"4", "3", "-"}
	// tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))
}
