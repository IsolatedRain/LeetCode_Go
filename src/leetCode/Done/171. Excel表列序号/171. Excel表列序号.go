package main

import "fmt"

func titleToNumber(s string) int {
	d := map[byte]int{}
	curByte := byte('A')
	for i := 1; i < 27; i++ {
		d[curByte] = i
		curByte++
	}
	res := 0
	for i := range s {
		res = res*26 + d[s[i]]
	}
	return res
}

func main() {
	s := "ZY"
	fmt.Println(titleToNumber(s))
}
