package main

import "fmt"

func hasAllCodes(s string, k int) bool {
	n := len(s)
	if n < k {
		return false
	}
	mark := make([]int, 1<<k)
	curMask := 0
	mask := (1 << k) - 1
	for i := 0; i < n; i++ {
		curMask <<= 1
		curMask |= int(s[i] - '0')
		if i >= k-1 {
			curMask &= mask
			mark[curMask] = 1
		}
	}
	for i := range mark {
		if mark[i] == 0 {
			return false
		}
	}
	return true
}

func main() {
	// s := "0000000001011100"
	// k := 4
	s := "00110110"
	k := 2
	r := hasAllCodes(s, k)
	fmt.Println(r)
}
