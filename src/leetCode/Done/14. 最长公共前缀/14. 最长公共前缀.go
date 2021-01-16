package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return strs[0]
	}
	left := longestCommonPrefix(strs[:n/2])
	right := longestCommonPrefix(strs[n/2:])
	return merge(left, right)
}
func merge(s1, s2 string) (res string) {
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	res = s1
	for i := range s1 {
		if s1[i] != s2[i] {
			res = s1[:i]
			break
		}
	}
	return
}

// func longestCommonPrefix(strs []string) string {
// 	n := len(strs)
// 	if n == 0 {
// 		return ""
// 	}
// 	if n == 1 {
// 		return strs[0]
// 	}
// 	sort.Slice(strs, func(i, j int) bool {
// 		s1, s2 := strs[i], strs[j]
// 		id := 0
// 		for id < len(s1) && id < len(s2) {
// 			if s1[id] < s2[id] {
// 				return true
// 			} else if s1[id] > s2[id] {
// 				return false
// 			}
// 			id++
// 		}
// 		if id == len(s1) && id == len(s2) {
// 			return true
// 		}
// 		if id < len(s1) {
// 			return false
// 		}
// 		return true
// 	})

// 	a, b := strs[0], strs[n-1]
// 	for i := range a {
// 		if a[i] != b[i] {
// 			return a[:i]
// 		}
// 	}
// 	return a
// }

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
