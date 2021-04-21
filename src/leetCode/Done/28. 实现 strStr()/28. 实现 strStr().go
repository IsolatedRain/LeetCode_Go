package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	n := len(haystack)
	if n == 0 {
		return -1
	}

	// 最长 真前缀==后缀 表
	preTable := make([]int, len(needle))
	// i: 前指针, 遍历字符串
	// j: 后指针, 指向当前相同长度的位置
	for i, j := 1, 0; i < len(needle); i++ {
		for j > 0 && needle[i] != needle[j] {
			fmt.Println(i, j, string(needle[i]), string(needle[j]))
			j = preTable[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		preTable[i] = j
	}
	fmt.Println(preTable)

	// i: haystack 指针
	// j: needle 指针
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = preTable[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}

	}
	return -1
}

func main() {
	// fmt.Println(strStr("aaaaaaab", "aaaaaab"))
	// fmt.Println(strStr("hello", "ll"))
	// fmt.Println(strStr("mississippi", "issip"))
	fmt.Println(strStr("aabc", "abcdabcc"))
}
