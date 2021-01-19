package main

import "fmt"

func firstUniqChar(s string) int {
	q := []int{}
	count := map[byte]int{}
	for i := range s {
		if _, ok := count[s[i]]; !ok {
			count[s[i]] = i
			q = append(q, i)
		} else {
			count[s[i]] = -1
			for len(q) > 0 && count[s[q[0]]] == -1 {
				q = q[1:]
			}
		}
	}
	if len(q) == 0 {
		return -1
	}
	return q[0]
}

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
}
