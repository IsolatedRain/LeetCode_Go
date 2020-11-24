package main

import "fmt"

func sortString(s string) string {
	letters := make([]int, 26)
	for i := range s {
		letters[s[i]-'a']++
	}
	res := make([]byte, 0, 26)
	for len(res) < len(s) {
		for i := 0; i < 26; i++ {
			if letters[i] > 0 {
				res = append(res, 'a'+byte(i))
				letters[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if letters[i] > 0 {
				res = append(res, 'a'+byte(i))
				letters[i]--
			}
		}
	}
	return string(res)
}

func main() {
	s := "aaaabbbbcccc"
	r := sortString(s)
	fmt.Printf("%v\n", r)
}
