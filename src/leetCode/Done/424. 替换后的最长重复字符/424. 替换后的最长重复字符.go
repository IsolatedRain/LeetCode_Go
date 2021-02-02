package main

import "fmt"

func characterReplacement(s string, k int) int {
	count := make([]int, 26)
	L := 0
	res := 0
	for i := range s {
		count[s[i]-'A']++
		for i-L+1-max(count) > k {
			count[s[L]-'A']--
			L++
		}
		res = max([]int{res, i - L + 1})
	}
	return res
}
func max(a []int) int {
	r := a[0]
	for _, v := range a {
		if v > r {
			r = v
		}
	}
	return r
}

func main() {
	s := "AABABBA"
	k := 2
	fmt.Println(characterReplacement(s, k))
}
