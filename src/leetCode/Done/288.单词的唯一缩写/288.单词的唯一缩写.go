package main

import (
	"fmt"
	"strconv"
)

// ValidWordAbbr 缩写类
type ValidWordAbbr struct {
	abbr map[string]map[string]struct{}
}

// Constructor 返回初始化后的缩写类
func Constructor(dictionary []string) ValidWordAbbr {
	res := make(map[string]map[string]struct{})
	for _, word := range dictionary {
		abbr := word
		if len(word) > 1 {
			abbr = word[:1] + strconv.Itoa(len(word)-2) + word[len(word)-1:]
		}
		if _, ok := res[abbr]; ok {
			res[abbr][word] = struct{}{}
		} else {
			res[abbr] = map[string]struct{}{word: {}}
		}
	}
	return ValidWordAbbr{res}
}

// IsUnique 检查是否唯一缩写
func (V *ValidWordAbbr) IsUnique(word string) bool {
	fmt.Println(V.abbr)
	curAbbr := word
	if len(word) > 1 {
		curAbbr = word[:1] + strconv.Itoa(len(word)-2) + word[len(word)-1:]
	}
	v, ok1 := V.abbr[curAbbr]
	if !ok1 {
		return true
	}
	if _, ok2 := V.abbr[curAbbr][word]; ok2 && len(v) == 1 {
		return true

	}
	return false
}

func main() {
	dictionary := []string{"deer", "door", "cake", "card"}
	obj := Constructor(dictionary)
	r := obj.IsUnique("cane")
	fmt.Println(r)
}

/**
 * Your ValidWordAbbr object will be instantiated and called as such:
 * obj := Constructor(dictionary);
 * param_1 := obj.IsUnique(word);
 */
