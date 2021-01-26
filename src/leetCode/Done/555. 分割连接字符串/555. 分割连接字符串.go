package main

import (
	"fmt"
	"strings"
)

func splitLoopedString(strs []string) string {
	s := []string{}
	for _, st := range strs {
		s = append(s, max(st, reversString(st)))
	}
	res := strings.Join(s, "")
	for i, word := range s {
		others := strings.Join(append(s[i+1:], s[:i]...), "")
		for j := range word {
			revWord := reversString(word)
			res = max(res, word[j:]+others+word[:j])
			res = max(res, revWord[j:]+others+revWord[:j])
		}
	}
	return res
}

func max(x, y string) string {
	if x > y {
		return x
	}
	return y
}
func reversString(s string) string {
	L, R, res := 0, len(s)-1, []byte(s)
	for L < R {
		res[L], res[R] = res[R], res[L]
		L++
		R--
	}
	return string(res)
}

func main() {
	fmt.Println(splitLoopedString([]string{"abc", "xyz", "efg"}))
}
