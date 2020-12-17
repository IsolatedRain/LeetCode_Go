package main

import "fmt"

func findTheDifference(s string, t string) byte {
	diff := t[len(t)-1]
	for i := range s {
		diff ^= s[i] ^ t[i]
	}
	return diff
}

func main() {
	s := "abcde"
	t := "abcfde"
	r := findTheDifference(s, t)
	fmt.Println(r)
}
