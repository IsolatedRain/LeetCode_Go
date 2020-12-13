package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	wordMap := map[string][]string{}
	for i := range strs {
		word := []byte(strs[i])
		sort.Slice(word, func(x, y int) bool {
			return word[x] < word[y]
		})
		k := string(word)
		wordMap[k] = append(wordMap[k], strs[i])
	}
	res := [][]string{}
	for _, v := range wordMap {
		res = append(res, v)
	}
	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// strs := []string{"ape", "and", "cat"}
	r := groupAnagrams(strs)
	fmt.Println(r)
}

// 计数数组当作键， 字典统计.
// func groupAnagrams(strs []string) [][]string {
// 	wordMap := map[[26]int][]string{}
// 	for _, word := range strs {
// 		count := [26]int{}
// 		for i := range word {
// 			count[word[i]-'a']++
// 		}
// 		wordMap[count] = append(wordMap[count], word)
// 	}
// 	res := [][]string{}
// 	for _, v := range wordMap {
// 		res = append(res, v)
// 	}
// 	return res
// }
