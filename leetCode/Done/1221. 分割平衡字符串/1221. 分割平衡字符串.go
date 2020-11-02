package main

import "fmt"

func balancedStringSplit(s string) int {
	d := map[rune]int{'R': 1, 'L': -1}
	curSum := 0
	res := 0
	for _, char := range s {
		curSum += d[char]
		if curSum == 0 {
			res++
		}
	}
	return res
}

func main() {
	s := "RLRRLLRLRL"
	r := balancedStringSplit(s)
	fmt.Printf("输入: %v\n输出: %v\n", s, r)
}
