package main

import "fmt"

func halvesAreAlike(s string) bool {
	vowels := map[byte]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}, 'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {}}
	n := len(s)
	cnt1, cnt2 := 0, 0
	for i := 0; i < n/2; i++ {
		if _, ok := vowels[s[i]]; ok {
			cnt1++
		}
		if _, ok := vowels[s[n-i-1]]; ok {
			cnt2++
		}
	}
	return cnt1 == cnt2
}

func main() {
	s := "AbCdEfGh"
	fmt.Println(halvesAreAlike(s))
}
