package main

import "fmt"

func permutation(s string) []string {
	n := len(s)
	curBytes := make([]byte, n)
	var dfs func(id, mask int)
	res := []string{}
	dfs = func(id, mask int) {
		if id == n {
			res = append(res, string(curBytes))
			return
		}
		for i := range s {
			if 1<<i&mask == 0 {
				curBytes[id] = s[i]
				dfs(id+1, mask|(1<<i))
			}
		}
	}

	dfs(0, 0)
	return res
}

func main() {
	fmt.Println(permutation("qwe"))
}
