package main

import (
	"fmt"
)

// 0, 1, 2, 3... numRows-1, numRows-2, numRows-3 ...3, 2, 1
// 模拟填入
func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}
	rows := make([][]byte, numRows)
	k := 0
	swap := 1
	for i := range s {
		rows[k] = append(rows[k], s[i])
		k += swap
		if k == numRows {
			swap = -1
			k -= 2
		} else if k == 0 {
			swap = 1
		}
	}
	res := ""
	for _, r := range rows {
		res += string(r)
	}
	return res
}

func main() {
	// s := "LEETCODEISHIRING"
	// numRows := 4
	s := "AB"
	numRows := 2
	r := convert(s, numRows)
	fmt.Println(r)
}
