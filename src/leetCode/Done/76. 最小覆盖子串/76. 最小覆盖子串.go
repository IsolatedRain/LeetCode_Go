package main

import "fmt"

func minWindow(s string, t string) string {
	countT := make(map[byte]int, 52)
	for i := range t {
		countT[t[i]]++
	}
	diff := len(countT)
	checked := map[byte]bool{}
	stack := []int{}
	cntDiffAlp := 0
	res := s
	countS := map[byte]int{}
	for i := range s {
		if _, ok := countT[s[i]]; !ok {
			continue
		}
		countS[s[i]]++
		if countS[s[i]] >= countT[s[i]] && !checked[s[i]] {
			checked[s[i]] = true
			cntDiffAlp++
		}
		stack = append(stack, i)
		for len(stack) > 0 && countS[s[stack[0]]] > countT[s[stack[0]]] {
			countS[s[stack[0]]]--
			stack = stack[1:]
		}
		if cntDiffAlp == diff {
			if i-stack[0] < len(res) {
				res = s[stack[0] : i+1]
			}
		}
	}
	if cntDiffAlp < diff {
		return ""
	}
	return res
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}
