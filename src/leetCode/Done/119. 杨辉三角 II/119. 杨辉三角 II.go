package main

import "fmt"

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1
	pre := 1
	for r := 1; r < rowIndex+1; r++ {
		res[r] = 1
		for i := 1; i < r; i++ {
			res[i], pre = pre+res[i], res[i]
		}
	}
	return res
}

func main() {
	fmt.Println(getRow(4))
}
