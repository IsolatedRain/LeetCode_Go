package main

import (
	"fmt"
	"sort"
)

func getKth(lo int, hi int, k int) int {
	var getW func(num int) int
	memo := map[int]int{}
	getW = func(num int) int {
		w := 0
		k := num
		for num != 1 {
			if v, ok := memo[num]; ok {
				w += v
				break
			}
			if num&1 == 1 {
				num = num*3 + 1
			} else {
				num /= 2
			}
			w++
		}
		memo[k] = w
		return w
	}
	arr := [][]int{}
	for i := lo; i <= hi; i++ {
		arr = append(arr, []int{getW(i), i})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0] || (arr[i][0] == arr[j][0] && arr[i][1] < arr[j][1])
	})
	return arr[k-1][1]
}

func main() {
	lo := 1
	hi := 1000
	k := 777
	fmt.Println(getKth(lo, hi, k))
}
