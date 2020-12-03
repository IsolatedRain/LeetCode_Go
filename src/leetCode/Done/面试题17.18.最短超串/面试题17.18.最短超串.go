package main

import "fmt"

func shortestSeq(big []int, small []int) []int {
	// small里数字 加入集合
	mapSet := make(map[int]struct{}, len(small))
	for _, v := range small {
		mapSet[v] = struct{}{}
	}

	var res []int
	var min func() int
	var L, curSize int
	minSize, index := len(big), map[int]int{}
	// 获取字典值里的 最小值（即最左索引）
	min = func() int {
		minID := len(big)
		for _, v := range index {
			if minID > v {
				minID = v
			}
		}
		return minID
	}

	for i, v := range big {
		if _, ok := mapSet[v]; ok {
			index[v] = i
		}
		if len(index) == len(small) && (res == nil || v == big[L]) {
			L = min()
			curSize = i - L
			if curSize < minSize {
				minSize = curSize
				res = []int{L, i}
			}
		}
	}
	return res
}

func main() {
	big := []int{7, 5, 9, 0, 2, 1, 3, 5, 7, 9, 1, 1, 5, 8, 8, 9, 7}
	small := []int{1, 5, 9}
	r := shortestSeq(big, small)
	fmt.Println(r)
}
