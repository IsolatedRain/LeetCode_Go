// https://leetcode-cn.com/problems/isomorphic-strings/
package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	a := map[byte]byte{}
	b := map[byte]byte{}
	for i := range s {
		v1, ok1 := a[s[i]]
		v2, ok2 := b[t[i]]
		if ok1 && v1 != t[i]{return false}
		if ok2 && v2 != s[i]{return false}
		a[s[i]] = t[i]
		b[t[i]] = s[i]
	}
	return true
}
func main() {
	fmt.Println(isIsomorphic("badc", "baba"))
}
