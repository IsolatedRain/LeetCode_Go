package main

import "fmt"

func isUgly(n int) bool {
	factors := []int{2, 3, 5}
	for _, fct := range factors {
		for n%fct == 0 {
			n /= fct
		}
	}
	return n == 1
}
func main() {
	fmt.Println(isUgly(8))
}
