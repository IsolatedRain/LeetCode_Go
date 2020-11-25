package main

import (
	"fmt"
	"strings"
)

func removeVowels(S string) string {
	for _, char := range "aoeiu" {
		S = strings.ReplaceAll(S, string(char), "")
	}
	// var tmp strings.Builder
	// for _, char := range S {
	// 	switch char {
	// 	case 'a', 'e', 'i', 'o', 'u':
	// 	default:
	// 		tmp.WriteRune(char)
	// 	}
	// }
	// return tmp.String()
	return S
}

func main() {
	s := "leetcodeisacommunityforcoders"
	r := removeVowels(s)
	fmt.Printf("输入: %v\n输出: %v\n", s, r)
}
