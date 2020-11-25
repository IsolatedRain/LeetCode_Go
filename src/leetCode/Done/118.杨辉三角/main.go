package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	n := numRows
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{1}}
	}
	arr := make([][]int, n+1)
	for i := 1; i < n+1; i++ {
		arr[i] = make([]int, i)
	}
	arr = arr[1:]
	arr[0][0] = 1
	arr[1] = []int{1, 1}
	for i := 2; i < n; i++ {
		arr[i][0] = 1
		arr[i][i] = 1
		for j := 1; j < i; j++ {
			arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
		}
	}
	// for _, a := range arr {
	// fmt.Printf("%v\n", a)
	// }
	return arr
}

func main() {
	p := 0
	r := generate(p)
	fmt.Printf("%v\n", r)
}
