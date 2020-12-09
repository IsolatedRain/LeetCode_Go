package main

import "fmt"

func uniquePaths(m int, n int) int {
	last := make([]int, m)
	for r := 0; r < n; r++ {
		cur := make([]int, m)
		for c := 0; c < m; c++ {
			if r == 0 || c == 0 {
				cur[c] = 1
			} else {
				cur[c] = cur[c-1] + last[c]
			}
		}
		last = cur
	}
	return last[m-1]
}

func main() {
	m := 3
	n := 7
	r := uniquePaths(m, n)
	fmt.Println(r)

}
