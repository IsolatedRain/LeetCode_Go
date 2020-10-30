package main

import (
	"fmt"
	"strings"
)

func addBinary(a string, b string) string {
	// 补齐长度, 使得len(a) == len(b)
	aLen, bLen := len(a), len(b)
	maxLen := max(aLen, bLen)
	var aDiff, bDiff strings.Builder
	for i := 0; i < maxLen-aLen; i++ {
		aDiff.WriteByte('0')
	}
	for i := 0; i < maxLen-bLen; i++ {
		bDiff.WriteByte('0')
	}
	a = aDiff.String() + a
	b = bDiff.String() + b

	// 双指针, 从后往前遍历
	aIdx := len(a) - 1
	bIdx := len(b) - 1
	carry := 0
	var tmp strings.Builder
	for aIdx > -1 && bIdx > -1 {
		// 同时为 1,  看进位 carry
		if a[aIdx] == '1' && b[bIdx] == '1' {
			if carry == 1 {
				tmp.WriteByte('1')
			} else {
				tmp.WriteByte('0')
				carry = 1
			}
			// 中之一为 1
		} else if a[aIdx] == '1' || b[bIdx] == '1' {
			if carry == 1 {
				tmp.WriteByte('0')
				carry = 1
			} else {
				tmp.WriteByte('1')
			}
			// 都为 0
		} else {
			if carry == 1 {
				tmp.WriteByte('1')
				carry = 0
			} else {
				tmp.WriteByte('0')
			}
		}
		aIdx--
		bIdx--
	}
	// 有进位 要加上
	if carry == 1 {
		tmp.WriteByte('1')
	}
	cur := tmp.String()
	// 翻转
	return reverse(cur)

}
func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func reverse(s string) string {
	res := []rune(s)
	for L, R := 0, len(s)-1; L < R; L, R = L+1, R-1 {
		res[L], res[R] = res[R], res[L]
	}
	return string(res)
}

func main() {
	// a := "111"
	a := "11"
	// b := "11"
	b := "10"
	r := addBinary(a, b)
	fmt.Printf("输入:%v  %v\n输出:%v", a, b, r)

}
