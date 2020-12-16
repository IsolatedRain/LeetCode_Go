package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	p2word := map[byte]string{}
	word2p := map[string]byte{}
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	for i := range pattern {
		_, ok1 := p2word[pattern[i]]
		_, ok2 := word2p[words[i]]
		if !ok1 && !ok2 {
			p2word[pattern[i]] = words[i]
			word2p[words[i]] = pattern[i]
		} else {
			if p2word[pattern[i]] != words[i] {
				return false
			}
			if word2p[words[i]] != pattern[i] {
				return false
			}
		}
	}
	return true
}

func main() {
	// pattern := "abba"
	// str := "dog cat cat dog"
	// pattern := "abba"
	// str := "dog dog dog dog"
	// pattern := "abba"
	// str := "dog cat cat dog"
	pattern := "abc"
	str := "dog cat dog"
	r := wordPattern(pattern, str)
	fmt.Println(r)
}
