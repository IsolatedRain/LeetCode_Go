package main

import (
	"fmt"
	"math"
)

func splitIntoFibonacci(S string) []int {
	fib := []int{}
	n := len(S)
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx == n {
			return len(fib) > 2
		}
		cur := 0
		for i := idx; i < n; i++ {
			if i > idx && S[idx] == '0' {
				return false
			}
			cur = cur*10 + int(S[i]-'0')
			if cur > math.MaxInt32 {
				return false
			}
			if len(fib) >= 2 {
				if cur < fib[len(fib)-1]+fib[len(fib)-2] {
					continue
				}
				if cur > fib[len(fib)-1]+fib[len(fib)-2] {
					return false
				}
			}
			fib = append(fib, cur)
			if dfs(i + 1) {
				return true
			}
			fib = fib[:len(fib)-1]
		}
		return false
	}
	dfs(0)
	return fib
}
func main() {
	S := "123456579"
	r := splitIntoFibonacci(S)
	fmt.Println(r)

}
