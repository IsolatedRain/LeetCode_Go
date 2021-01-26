package main

import (
	"fmt"
	"strconv"
)

func generateAbbreviations(word string) []string {
	n := len(word)
	tar := 1 << n
	res := []string{}
	for i := 0; i < tar; i++ {
		curCnt := 0
		curWord := ""
		for j := 0; j < n; j++ {
			if 1<<j&i == 0 {
				curCnt++
			} else {
				if curCnt != 0 {
					curWord += strconv.Itoa(curCnt) + string(word[j])
				} else {
					curWord += string(word[j])
				}
				curCnt = 0
			}
		}
		if curCnt != 0 {
			curWord += strconv.Itoa(curCnt)
		}
		res = append(res, curWord)

	}
	return res
}

func main() {
	fmt.Println(generateAbbreviations("word"))
}
