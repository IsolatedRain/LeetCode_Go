package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	d := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'}}
	res := []string{}
	n := len(digits)
	curBytes := make([]byte, n, n)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			res = append(res, string(curBytes))
			return
		}
		for _, b := range d[digits[i]] {
			curBytes[i] = b
			dfs(i + 1)
		}
		return
	}
	dfs(0)
	return res
}

func main() {
	digits := "23"
	r := letterCombinations(digits)
	fmt.Println(r)

}
