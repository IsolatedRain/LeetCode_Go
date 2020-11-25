package main

import "fmt"

func generatePossibleNextMoves(s string) []string {
	n := len(s)
	res := []string{}
	for i := 0; i < n-1; i++ {
		if s[i:i+2] == "++" {
			res = append(res, s[:i]+"++"+s[i+2:])
		}
	}
	// for i := 0; i < n; i++ {
	// 	if i >= n-1 || s[i] != '+' || s[i] != s[i+1] {
	// 		continue
	// 	}
	// 	curBytes := []byte(s)
	// 	curBytes[i] = '-'
	// 	curBytes[i+1] = '-'
	// 	res = append(res, string(curBytes))
	// }
	return res
}

func main() {
	s := "+-+-++--+-+"
	r := generatePossibleNextMoves(s)
	fmt.Printf("输入: %v\n输出: %v\n", s, r)
}
