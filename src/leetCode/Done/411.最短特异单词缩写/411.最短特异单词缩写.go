package main

import (
	"fmt"
	"sort"
	"strconv"
)

func minAbbreviation(target string, dictionary []string) string {
	tarLen := len(target)
	sameLen := []string{}
	for _, word := range dictionary {
		if len(word) == tarLen {
			sameLen = append(sameLen, word)
		}
	}
	if len(sameLen) == 0 {
		return strconv.Itoa(tarLen)
	}
	targetAbbrs := getAllAbbrs(target)
	tarAbbrsLen := map[string]int{}
	for _, abbr := range targetAbbrs {
		tarAbbrsLen[abbr] = calLen(abbr)
	}
	sort.Slice(targetAbbrs, func(i, j int) bool {
		return tarAbbrsLen[targetAbbrs[i]] < tarAbbrsLen[targetAbbrs[j]]
	})

	for _, abbr := range targetAbbrs {
		i := 0
		for i < len(sameLen) {
			if isAbbr(sameLen[i], abbr) {
				break
			}
			i++
		}
		if i == len(sameLen) {
			return abbr
		}
	}
	return ""
}
func isAbbr(word, abbr string) bool {
	wordLen := len(word)
	wordID := 0
	curNum := 0
	for _, b := range []byte(abbr) {
		if isDigit(b) {
			curNum = curNum*10 + int(b-'0')
		} else {
			wordID += curNum
			curNum = 0
			if word[wordID] != b {
				return false
			}
			wordID++
		}
	}
	return wordID+curNum == wordLen
}
func isDigit(s byte) bool {
	if '0' <= s && s <= '9' {
		return true
	}
	return false
}
func calLen(s string) int {
	curLen := 0
	curNum := 0
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			curNum = 1
		} else {
			curLen += curNum + 1
			curNum = 0
		}
	}
	return curLen + curNum
}
func getAllAbbrs(word string) []string {
	n := len(word)
	res := []string{}
	var dfs func(i, k int, path string)
	dfs = func(i, k int, path string) {
		if i == n {
			if k != 0 {
				path += strconv.Itoa(k)
			}
			res = append(res, path)
			return
		}
		if k == 0 {
			dfs(i+1, 0, path+string(word[i]))
		} else {
			dfs(i+1, 0, path+strconv.Itoa(k)+string(word[i]))
		}
		dfs(i+1, k+1, path)
	}
	dfs(0, 0, "")
	return res
}

func main() {
	// target := "apple"
	// dictionary := []string{"blade"}
	// dictionary := []string{"kkkk", "blade"}
	target := "aaaaaxaaaaa"
	dictionary := []string{"bbbbbxbbbbb"}
	r := minAbbreviation(target, dictionary)
	fmt.Println(r)
}
