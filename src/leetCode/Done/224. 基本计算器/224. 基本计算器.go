package main

import (
	"fmt"
)

func calculate(s string) (ans int) {
	ops := []int{1}
	sign := 1
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = ops[len(ops)-1]
			i++
		case '-':
			sign = -ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			ans += sign * num
		}
	}
	return
}
func cal(op string, a, b int) int {
	if op == "+" {
		return a + b
	}
	return a - b
}
func isDigit(b byte) bool {
	if '0' <= b && b <= '9' {
		return true
	}
	return false
}

func main() {
	// s := "(1+( 4 + 5+2)-3)+(6 + 8) + 12"
	s := "-2+ 1"
	fmt.Println(calculate(s))
}
