//https://www.codewars.com/kata/56f3a1e899b386da78000732/train/go
package main

import (
	"fmt"
	"strings"
)

// PartList 把列表分为 非空的 两部分
func PartList(arr []string) string {
	s := ""
	for i := 1; i < len(arr); i++ {
		left := strings.Join(arr[:i], " ")
		right := strings.Join(arr[i:], " ")
		s += fmt.Sprintf("(%s, %s)", left, right)
	}
	return s
}

func main() {
	fmt.Println(PartList([]string{"I", "wish", "I", "hadn't", "come"}))
}
