package main

import "fmt"

// 双向BFS
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordsSet := map[string]int{}
	for i, w := range wordList {
		wordsSet[w] = i
	}

	startWords := map[string]struct{}{beginWord: {}}
	startUsed := make([]bool, len(wordList))

	endWords := map[string]struct{}{endWord: {}}
	endUsed := make([]bool, len(wordList))

	if i, ok := wordsSet[endWord]; ok {
		endUsed[i] = true
	} else {
		return 0
	}

	steps := 0
	for len(startWords) > 0 {
		steps++
		nextStart := map[string]struct{}{}

		for word := range startWords {
			chars := []byte(word)
			for i := 0; i < len(word); i++ {
				swap := chars[i]
				for c := 'a'; c <= 'z'; c++ {
					chars[i] = byte(c)
					newWord := string(chars)
					idx, ok := wordsSet[newWord]
					if !ok || startUsed[idx] {
						continue
					}
					if endUsed[idx] {
						return steps + 1
					}
					nextStart[newWord] = struct{}{}
					startUsed[idx] = true
				}
				chars[i] = swap
			}
		}

		startWords = nextStart
		if len(startWords) > len(endWords) {
			startWords, endWords = endWords, startWords
			startUsed, endUsed = endUsed, startUsed
		}
	}

	return 0
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Printf("输入:	%v\n	%v\n	%v\n", beginWord, endWord, wordList)
	r := ladderLength(beginWord, endWord, wordList)
	fmt.Printf("输出:	%v\n", r)
}
