package main

import "fmt"

func isOneEditDistance(s string, t string) bool {
	if len(s) > len(t) {
		s, t = t, s
	}
	sLen := len(s)
	tLen := len(t)
	if tLen-sLen > 1 {
		return false
	}
	if sLen == 0 || tLen == 0 {
		return t != s
	}
	if tLen == sLen {
		diff := 0
		for i := range s {
			if s[i] != t[i] {
				diff++
			}
			if diff > 1 {
				return false
			}
		}
	} else {
		for i := range s {
			if s[i] != t[i] {
				return s[i:] == t[i+1:]
			}
		}
	}
	return true
}

func main() {
	s := "123"
	t := "1203"
	r := isOneEditDistance(s, t)
	fmt.Println(r)
}
