package main

import "fmt"

func isHappy(n int) bool {
	checked := map[int]bool{}
	var cal func(x int) int
	cal = func(x int) int {
		r := 0
		for x > 0 {
			a := x % 10
			x /= 10
			r += a * a
		}
		return r
	}
	for {
		if n == 1 {
			return true
		}
		n = cal(n)
		if checked[n] == true {
			break
		}
		checked[n] = true
	}
	return false
}

func main() {
	fmt.Println(isHappy(19))
}
