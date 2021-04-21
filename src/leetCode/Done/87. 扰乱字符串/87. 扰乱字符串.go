package main

import "fmt"

func isScramble(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 != n2 {
		return false
	}
	if !count(s1, s2) {
		return false
	}
	var f func(a, b string) bool
	memo := map[string]bool{}
	f = func(left, right string) bool {
		k := left + " " + right
		if v, ok := memo[k]; ok {
			return v
		}

		curRes := false
		defer func() { memo[k] = curRes }()
		if left == right {
			curRes = true
			return true
		}
		if !count(left, right) {
			return false
		}
		curLen := len(left)
		for i := 1; i < curLen; i++ {
			if f(left[:i], right[:i]) && f(left[i:], right[i:]) {
				curRes = true
				return true
			}
			if f(left[:i], right[curLen-i:]) && f(left[i:], right[:curLen-i]) {
				curRes = true
				return true
			}
		}
		return false
	}
	ret := f(s1, s2)
	return ret
}
func count(s1, s2 string) bool {
	a := [26]int{}
	b := [26]int{}
	for i := range s1 {
		a[s1[i]-'a']++
	}
	for i := range s2 {
		b[s2[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	// s1 := "great"
	// s2 := "rgeat"
	s1 := "eebaacbcbcadaaedceaaacadccd"
	s2 := "eadcaacabaddaceacbceaabeccd"
	fmt.Println(isScramble(s1, s2))
}
