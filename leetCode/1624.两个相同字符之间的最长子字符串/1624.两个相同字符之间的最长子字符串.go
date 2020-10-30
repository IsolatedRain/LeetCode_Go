package main

import (
	"fmt"
)

// 记录第一次出现的索引
func maxLengthBetweenEqualCharacters(s string) int {
	d := make(map[rune]int, len(s))
	res := -1
	for i, c := range s {
		if leftIdx, exist := d[c]; !exist {
			d[c] = i
		} else {
			res = max(res, i-leftIdx-1)
		}
	}
	return res
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	p := "cabbac"
	fmt.Printf("输入: %v\n", p)
	r := maxLengthBetweenEqualCharacters(p)
	fmt.Printf("输出: %v\n", r)
}
