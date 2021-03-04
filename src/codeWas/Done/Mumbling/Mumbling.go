// https://www.codewars.com/kata/5667e8f4e3f572a8f2000039/train/go
package main

import (
	"fmt"
	"strings"
)

// Accum 构造新字符串
func Accum(s string) string {
	res := ""
	s1 := strings.ToLower(s)
	for i, c := range s1 {
		s2 := string(c)
		head := strings.ToUpper(s2)
		tail := strings.Repeat(strings.ToLower(s2), i)
		res += fmt.Sprintf("%s%s-", head, tail)
	}
	return res[:len(res)-1]
}

func main() {
	fmt.Println(Accum("ZpglnRxqenU") == "Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu")
}
