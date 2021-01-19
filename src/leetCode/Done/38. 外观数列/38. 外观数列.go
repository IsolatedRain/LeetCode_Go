package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	var f func(s string) string
	f = func(s string) string {
		s += " "
		chars := []string{}
		L := 0
		for i := 1; i < len(s); i++ {
			if s[i] != s[i-1] {
				chars = append(chars, s[L:i])
				L = i
			}
		}
		r := ""
		for _, c := range chars {
			r += strconv.Itoa(len(c)) + c[:1]
		}
		return r
	}

	res := "1"
	for n > 1 {
		n--
		res = f(res)
	}
	return res
}

func main() {
	fmt.Println(countAndSay(7))
}
