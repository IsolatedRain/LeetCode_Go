package main

import "fmt"

func letterCasePermutation(S string) []string {
	countAlp := 0
	for i := range S {
		if isAlpha(S[i]) {
			countAlp++
		}
	}
	maxMask := 1 << countAlp
	res := []string{}
	for mask := 0; mask < maxMask; mask++ {
		curWord := []byte(S)
		alpID := 0
		for i := range S {
			if !isAlpha(S[i]) {
				continue
			}
			if 1<<alpID&mask > 0 {
				curWord[i] = upper(S[i])
			} else {
				curWord[i] = lower(S[i])
			}
			alpID++
		}
		res = append(res, string(curWord))
	}
	return res
}
func lower(b byte) byte {
	if 'a' <= b && b <= 'z' {
		return b
	}
	return b + 32
}
func upper(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b
	}
	return b - 32
}
func isAlpha(b byte) bool {
	if ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') {
		return true
	}
	return false
}

func main() {
	S := "a1b2"
	fmt.Println(letterCasePermutation(S))
}
