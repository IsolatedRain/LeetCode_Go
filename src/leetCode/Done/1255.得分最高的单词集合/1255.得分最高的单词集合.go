package main

import "fmt"

func maxScoreWords(words []string, letters []byte, score []int) int {
	count := map[byte]int{}
	for _, b := range letters {
		count[b]++
	}
	validWords := []string{}
	for _, word := range words {
		curCnt := map[byte]int{}
		for i := range word {
			curCnt[word[i]]++
			if v, ok := count[word[i]]; !ok || curCnt[word[i]] > v {
				break
			}
		}
		validWords = append(validWords, word)
	}
	n := len(validWords)
	var calScore func(word string) int
	calScore = func(word string) (res int) {
		for i := range word {
			res += score[word[i]-'a']

		}
		return
	}
	res := 0
	var dfs func(i, curScore int, cnt []int)
	dfs = func(i, curScore int, cnt []int) {
		if i == n {
			res = max(res, curScore)
			return
		}
		dfs(i+1, curScore, cnt)
		curCnt := make([]int, 26)
		copy(curCnt, cnt)
		for j := range validWords[i] {
			curCnt[validWords[i][j]-'a']++
			if curCnt[validWords[i][j]-'a'] > count[validWords[i][j]] {
				return
			}
		}
		curScore += calScore(validWords[i])
		dfs(i+1, curScore, curCnt)
		return
	}
	dfs(0, 0, make([]int, 26))
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	words := []string{"dog", "cat", "dad", "good"}
	letters := []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'}
	score := []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	r := maxScoreWords(words, letters, score)
	fmt.Println(r)
}

// 状态压缩
// func maxScoreWords(words []string, letters []byte, score []int) int {
// 	count := map[byte]int{}
// 	for _, b := range letters {
// 		count[b]++
// 	}
// 	validWords := []string{}
// 	for _, word := range words {
// 		curCnt := map[byte]int{}
// 		for i := range word {
// 			curCnt[word[i]]++
// 			if v, ok := count[word[i]]; !ok || curCnt[word[i]] > v {
// 				break
// 			}
// 		}
// 		validWords = append(validWords, word)
// 	}
// 	n := len(validWords)
// 	maxMask := 1 << n
// 	res := 0
// 	for m := 1; m < maxMask; m++ {
// 		curScore := 0
// 		curCnt := map[byte]int{}
// 		valid := true
// 		for i := range validWords {
// 			if (m>>i)&1 == 1 {
// 				curWord := validWords[i]
// 				for j := range curWord {
// 					curCnt[curWord[j]]++
// 					curScore += score[curWord[j]-'a']
// 					if curCnt[curWord[j]] > count[curWord[j]] {
// 						valid = false
// 						break
// 					}
// 				}
// 			}
// 			if !valid {
// 				break
// 			}
// 		}
// 		if valid {
// 			res = max(res, curScore)
// 		}
// 	}
// 	return res
// }
