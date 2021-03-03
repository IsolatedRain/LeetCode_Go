// https://www.codewars.com/kata/55e6f5e58f7817808e00002e/solutions/go
package main

import "fmt"

// Seven 被 7 整除
func Seven(n int64) []int {
	res := []int{}
	var dfs func(x int64, steps int)
	dfs = func(x int64, steps int) {
		if x < 100 {
			res = append(res, []int{int(x), steps}...)
			return
		}
		left := x / 10
		right := x % 10
		nextX := left - 2*right
		dfs(nextX, steps+1)
		return
	}
	dfs(n, 0)
	return res
}

func main() {
	fmt.Println(Seven(477557102), []int{47, 7})

}
