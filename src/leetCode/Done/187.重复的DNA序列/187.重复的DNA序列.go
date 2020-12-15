package main

import "fmt"

// 状压， 每次左移两位，加入新字母， 去掉最右2位。
func findRepeatedDnaSequences(s string) []string {
	n := len(s)
	if n <= 10 {
		return []string{}
	}
	mask := (1 << 20) - 1
	toInt := map[byte]int{'A': 0, 'T': 1, 'C': 2, 'G': 3}
	bitMask := 0
	seen := map[int]string{}
	res := map[string]struct{}{}
	for i := 0; i < n-9; i++ {
		if i == 0 {
			for j := 0; j < 10; j++ {
				bitMask <<= 2
				bitMask |= toInt[s[j]]
			}
		} else {
			bitMask <<= 2
			bitMask |= toInt[s[i+10-1]]
			bitMask &= mask
		}
		if _, ok := seen[bitMask]; ok {
			res[s[i:i+10]] = struct{}{}
		}
		seen[bitMask] = s[i : i+10]
	}
	output := []string{}
	for k := range res {
		output = append(output, k)
	}
	return output
}

func main() {
	// s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	// s := "AAAAAAAAAAAAA"
	s := "AAAAAAAAAAA"
	r := findRepeatedDnaSequences(s)
	fmt.Println(r)
}
