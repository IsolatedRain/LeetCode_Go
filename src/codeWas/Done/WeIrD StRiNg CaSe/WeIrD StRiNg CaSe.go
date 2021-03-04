// https://www.codewars.com/kata/52b757663a95b11b3d00062d/train/go
package main

import (
	"fmt"
	"strings"
)

func toWeirdCase(s string) string {
	help := func(word string) string {
		r := ""
		for i := range word {
			if i&1 == 1 {
				r += strings.ToLower(string(word[i]))
			} else {
				r += strings.ToUpper(string(word[i]))
			}
		}
		return r
	}

	words := strings.Split(s, " ")
	res := []string{}
	for _, word := range words {
		res = append(res, help(word))
	}
	return strings.Join(res, " ")
}

func main() {
	fmt.Println(toWeirdCase("This is a test Looks like you passed") == "ThIs Is A TeSt LoOkS LiKe YoU PaSsEd")
}
