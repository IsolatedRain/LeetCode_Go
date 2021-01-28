package main

import (
	"fmt"
	"sort"
	"strings"
)

func beforeAndAfterPuzzles(phrases []string) []string {
	res := []string{}
	words := [][]string{}
	for _, p := range phrases {
		words = append(words, strings.Split(p, " "))
	}
	set := map[string]struct{}{}
	for i := range words {
		for j := range words {
			if i != j && words[i][len(words[i])-1] == words[j][0] {
				set[strings.Join(append(words[i], words[j][1:]...), " ")] = struct{}{}
			}
		}
	}
	for k := range set {
		res = append(res, k)
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

func main() {
	fmt.Println(beforeAndAfterPuzzles([]string{"writing code", "code rocks"}))
}
