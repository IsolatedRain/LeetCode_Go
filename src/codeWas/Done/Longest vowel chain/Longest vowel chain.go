// https://www.codewars.com/kata/59c5f4e9d751df43cf000035/train/go
package main

import "fmt"

// Solve 最大的连续元音长度
func Solve(s string) int {
	s += " "
	res, L, R, n := 0, 0, 0, len(s)
	mapSet := map[byte]bool{'o': true, 'e': true, 'a': true, 'i': true, 'u': true}
	for R < n {
		if _, ok := mapSet[s[R]]; !ok {
			L = R + 1
		} else {
			fmt.Println(L, R)
			res = max(res, R-L+1)
		}
		R++
	}
	return res
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(Solve("codewarriors"))
}
