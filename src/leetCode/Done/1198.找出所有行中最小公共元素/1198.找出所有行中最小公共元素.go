package main

import (
	"fmt"
	_ "sort"
)

func smallestCommonElement(mat [][]int) int {
	n := len(mat)
	m := len(mat[0])
	// 每一行的当前索引
	ids := make([]int, n, n)
	// 当前索引 至少 >= 的数
	curVal := mat[0][0]
	for {
		tmp := curVal
		for i := 0; i < n; i++ {
			for mat[i][ids[i]] < curVal {
				ids[i]++
				if ids[i] >= m {
					return -1
				}
			}
			curVal = max(curVal, mat[i][ids[i]])
		}
		if curVal == tmp {
			break
		}
	}
	return curVal
}

// n := len(mat)
// matCnt := map[int]int{}
// for r := range mat {
// 	for c := range mat[r] {
// 		matCnt[mat[r][c]]++
// 	}
// }
// minCommon := []int{}
// for k, v := range matCnt {
// 	if v == n {
// 		minCommon = append(minCommon, k)
// 	}
// }
// if len(minCommon) < 1 {
// 	return -1
// }
// sort.Slice(minCommon, func(i, j int) bool {
// 	return minCommon[i] < minCommon[j]
// })
// return minCommon[0]
// }
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	mat := [][]int{{1, 2, 3, 4, 5}, {2, 4, 6, 8, 10}, {3, 5, 7, 9, 11}, {1, 3, 5, 7, 9}}
	r := smallestCommonElement(mat)
	fmt.Println(r)
}
