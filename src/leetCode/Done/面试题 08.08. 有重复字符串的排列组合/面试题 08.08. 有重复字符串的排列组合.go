package main

import (
	"fmt"
	"sort"
)

func permutation(S string) []string {
	s := []byte(S)
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	res, n := []string{}, len(s)
	curBytes := make([]byte, n)
	var dfs func(id, mask int)
	dfs = func(id, mask int) {
		if id == n {
			res = append(res, string(curBytes))
			return
		}
		mark := map[byte]bool{}
		for i := range s {
			if 1<<i&mask == 0 && !mark[s[i]] {
				mark[s[i]] = true
				curBytes[id] = s[i]
				dfs(id+1, mask|(1<<i))
			}
		}
	}
	dfs(0, 0)
	return res
}

func main() {
	fmt.Println(permutation("qqe"))
}
