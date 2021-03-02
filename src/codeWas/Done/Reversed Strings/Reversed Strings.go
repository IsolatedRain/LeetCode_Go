package main

import "fmt"

// Solution 翻转字符串
func Solution(word string) string {
	b := []byte(word)
	L, R := 0, len(b)-1
	for L < R {
		b[L], b[R] = b[R], b[L]
		L++
		R--
	}
	return string(b)
}
func main() {
	fmt.Println(Solution("word"))
}
