package main

import "fmt"

func clumsy(N int) int {
	stack := []int{N}
	N = N - 1
	i := 0
	sign := []string{"*", "/", "+", "-"}
	for ; N > 0; N-- {
		op := sign[i%4]
		switch op {
		case "-":
			stack = append(stack, -N)
		default:
			stack[len(stack)-1] = cal(stack[len(stack)-1], N, op)
		}
		i++
	}
	return sum(stack...)
}
func cal(x, y int, op string) int {
	switch op {
	case "*":
		return x * y
	case "/":
		return x / y
	default:
		return x + y
	}
}
func sum(x ...int) int {
	if len(x) == 0 {
		return 0
	}
	v := x[0]
	for i := 1; i < len(x); i++ {
		v += x[i]
	}
	return v
}

func main() {
	fmt.Println(clumsy(10))
}
