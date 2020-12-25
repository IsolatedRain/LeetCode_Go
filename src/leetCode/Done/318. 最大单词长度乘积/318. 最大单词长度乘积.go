package main

import "fmt"

func maxProduct(words []string) int {
	wordsMap := map[int]int{}
	for _, word := range words {
		curMask := 0
		for i := range word {
			curMask |= 1 << (word[i] - 'a')
		}
		wordsMap[curMask] = max(wordsMap[curMask], len(word))
	}
	w := []int{}
	for k := range wordsMap {
		w = append(w, k)
	}
	n := len(w)
	res := 0
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if w[i]&w[j] == 0 {
				res = max(res, wordsMap[w[i]]*wordsMap[w[j]])
			}
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	words := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	r := maxProduct(words)
	fmt.Println(r)
}
