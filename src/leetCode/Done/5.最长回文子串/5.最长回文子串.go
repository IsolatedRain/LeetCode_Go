package main

import "fmt"

func longestPalindrome(s string) string {
	n := len(s)
	arr := make([][]bool, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]bool, n)
	}
	res := ""
	curLen := 0
	for i := 0; i < n; i++ {
		arr[i][i] = true
		for j := 0; j < i; j++ {
			if s[i] == s[j] {
				if (i-j) == 1 || arr[i-1][j+1] == true {
					arr[i][j] = true
					if i-j+1 > curLen {
						curLen = i - j + 1
						res = s[j : i+1]
					}
				}
			}
		}
	}

	if curLen == 0 {
		res = s[:1]
	}
	return res
}
func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	// fmt.Printf("%v\n", longestPalindrome("babad"))
	fmt.Printf("%v\n", longestPalindrome("cbdbc"))
}
