// https://www.codewars.com/kata/57eb8fcdf670e99d9b000272/train/go
package main

import (
	"fmt"
	"strings"
)

// High 返回 最大字母组合
func High(s string) string {
	calScore := func(s string) int {
		r := 0
		for i := range s {
			r += int(s[i]-'a') + 1
		}
		return r
	}
	words := strings.Split(s, " ")

	res, maxScore := "", 0
	for _, word := range words {
		curScore := calScore(word)
		if curScore > maxScore {
			res = word
			maxScore = curScore
		}
	}
	return res
}

func main() {
	// fmt.Println(High("what time are we climbing up the volcano") == "volcano")
	fmt.Println(High("d bb"))
}
