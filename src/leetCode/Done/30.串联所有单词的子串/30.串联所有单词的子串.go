package main

import "fmt"

func findSubstring(s string, words []string) []int {
	subLen := len(words[0])
	wordsMap := make(map[string]int)
	for _, w := range words {
		wordsMap[w]++
	}
	match := len(wordsMap)
	res := []int{}
	for i := 0; i < subLen; i++ {
		L := i
		curMatch := 0
		curMap := map[string]int{}
		// 以subLen 为间隔，遍历。
		for j := i; j <= len(s)-subLen; j += subLen {
			curWord := s[j : j+subLen]
			_, ok := wordsMap[curWord]
			// 如果不在字典里， 则移动左边界至当前位置+subLen
			if !ok {
				L = j + subLen
				curMatch = 0
				curMap = map[string]int{}
				continue
			}
			if ok {
				curMap[curWord]++
			}
			// 统计和单词字典里个数一样的单词的计数。
			// 如果字典里的单词一样且其计数一样， 则两字典相同，即符合串联所有单词。
			if curMap[curWord] == wordsMap[curWord] {
				curMatch++
			}
			if curMatch == match {
				res = append(res, L)
				curMatch--
				curMap[s[L:L+subLen]]--
				L += subLen
			}
			// 如果超过了字典里的计数， 则以subLen为step 移动左边界，
			// 同时减去计数。
			for curMap[curWord] > wordsMap[curWord] {
				word := s[L : L+subLen]
				if curMap[word] == wordsMap[word] {
					curMatch--
				}
				curMap[word]--
				L += subLen
			}
		}
	}
	return res
}
func main() {
	// s := "barfoothefoobarman"
	// words := []string{"foo", "bar"}
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	r := findSubstring(s, words)
	fmt.Println(r)

}
