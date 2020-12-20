package main

import "fmt"

// 逐行更新的同时，记录最小的两个数，以及最小数的列索引。
func minFallingPathSum(arr [][]int) int {
	row := len(arr)
	if row == 1 {
		return min(arr[0]...)
	}
	col := len(arr[0])
	lastMin := 100
	lastMinID := -1
	lastSecMin := 100
	for i, v := range arr[0] {
		if v < lastMin {
			lastMin, lastSecMin = v, lastMin
			lastMinID = i
		} else if v < lastSecMin {
			lastSecMin = v
		}
	}

	for i := 1; i < row; i++ {
		curMin, curSecMin := 100000, 100000
		curMinID := -1
		for j := 0; j < col; j++ {
			curV := arr[i][j]
			if j == lastMinID {
				curV += lastSecMin
			} else {
				curV += lastMin
			}
			if curV < curMin {
				curMin, curSecMin = curV, curMin
				curMinID = j
			} else if curV < curSecMin {
				curSecMin = curV
			}
		}
		lastMin, lastSecMin, lastMinID = curMin, curSecMin, curMinID
	}
	return lastMin
}

func min(x ...int) int {
	r := x[0]
	for _, v := range x {
		if r > v {
			r = v
		}
	}
	return r
}

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	r := minFallingPathSum(arr)
	fmt.Println(r)
}
