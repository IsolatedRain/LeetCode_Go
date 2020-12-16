package main

import "fmt"

func findNumOfValidWords(words []string, puzzles []string) []int {
	var getMask func(w string) int
	getMask = func(w string) int {
		wordMask := 0
		for i := range w {
			wordMask |= 1 << (w[i] - 'a')
		}
		return wordMask
	}
	maskCount := map[int]int{}
	for _, word := range words {
		curK := getMask(word)
		maskCount[curK]++
	}
	res := []int{}
	for _, puzzle := range puzzles {
		curMask := getMask(puzzle)
		subMask := curMask
		headMask := 1 << int(puzzle[0]-'a')
		curRes := 0
		for subMask > 0 {
			if c, ok := maskCount[subMask]; ok && subMask&headMask > 0 {
				curRes += c
			}
			subMask = curMask & (subMask - 1)
		}
		res = append(res, curRes)
	}
	return res
}

func main() {
	words := []string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"}
	puzzles := []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}
	r := findNumOfValidWords(words, puzzles)
	fmt.Println(r)
}
