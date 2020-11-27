package main

import (
	"fmt"
	"sort"
)

func smallestCommonElement(mat [][]int) int {
	n := len(mat)
	matCnt := map[int]int{}
	for r := range mat {
		for c := range mat[r] {
			matCnt[mat[r][c]]++
		}
	}
	minCommon := []int{}
	for k, v := range matCnt {
		if v == n {
			minCommon = append(minCommon, k)
		}
	}
	if len(minCommon) < 1 {
		return -1
	}
	sort.Slice(minCommon, func(i, j int) bool {
		return minCommon[i] < minCommon[j]
	})
	return minCommon[0]
}

func main() {
	mat := [][]int{{1, 2, 3, 4, 5}, {2, 4, 5, 8, 10}, {3, 5, 7, 9, 11}, {1, 3, 5, 7, 9}}
	r := smallestCommonElement(mat)
	fmt.Println(r)
}
