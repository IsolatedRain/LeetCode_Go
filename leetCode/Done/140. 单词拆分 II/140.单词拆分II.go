package main

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	wordSet := map[string]struct{}{}
	for _, w := range wordDict {
		wordSet[w] = struct{}{}
	}

	sLen := len(s)
	dp := make([][][]string, sLen)
	var dfs func(i int) [][]string
	dfs = func(idx int) [][]string {
		if idx == sLen {
			return nil
		}
		if dp[idx] != nil {
			return dp[idx]
		}
		wordList := [][]string{}
		for i := idx; i < sLen; i++ {
			word := s[idx : i+1]
			if _, ok := wordSet[word]; ok {
				nextWordsList := dfs(i + 1)
				curList := []string{word}
				if nextWordsList == nil {
					wordList = append(wordList, curList)
				} else {
					for _, nextWords := range nextWordsList {
						wordList = append(wordList, append([]string{word}, nextWords...))
					}
				}
			}
		}
		dp[idx] = wordList
		return wordList
	}
	words := dfs(0)
	res := make([]string, 0, len(words))
	fmt.Printf("%v\n", dp)
	for _, word := range words {
		res = append(res, strings.Join(word, " "))
	}

	return res
}

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Printf("输入: %v %v\n", s, wordDict)
	res := wordBreak(s, wordDict)
	fmt.Printf("输出: %v\n", res)
}
