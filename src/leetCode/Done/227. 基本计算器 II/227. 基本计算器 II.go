package main

import "fmt"

func calculate(s string) (res int) {
	sign := byte('+')
	n := len(s)
	stack := []int{}
	num := 0
	for i := 0; i < n; {
		isDigit := '0' <= s[i] && s[i] <= '9'
		if isDigit {
			num = num*10 + int(s[i]-'0')
		}
		if !isDigit && s[i] != ' ' || i == n-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			sign = s[i]
			num = 0
		}
		i++
	}
	for _, num := range stack {
		res += num
	}
	return
}

func main() {
	fmt.Println(calculate(" 3+5 / 2 "))
}
