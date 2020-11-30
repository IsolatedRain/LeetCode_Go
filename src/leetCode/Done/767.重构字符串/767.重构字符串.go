package main

import (
	"fmt"
	"sort"
)
// 最多的字符个数  不能超过1半
// 按个数排序，双指针，头+中间， 逐对添加。
func reorganizeString(S string) string {
	n := len(S)
	count := make(map[byte]int, n)
	for i := range S {
		count[S[i]]++
	}
	k := [][]int{}
	for ky, v := range count {
		k = append(k, []int{int(ky), v})
	}
	sort.Slice(k, func(i, j int) bool {
		return k[i][1] > k[j][1]
	})
	if n&1 == 1 {
		if k[0][1] > (n>>1)+1 {
			return ""
		}
	}else{
		if k[0][1] > n >> 1{
			return ""
		}
	}
	

	s := []byte{}
	for i := range k {
		for j := k[i][1]; j > 0; j-- {
			s = append(s, byte(k[i][0]))
		}
	}
	L := 0
	R := n >> 1
	if n&1 == 1 {
		R++
	}
	res := []byte{}
	for R < n {
		res = append(res, s[L], s[R])
		R++
		L++
	}
	if n&1 == 1 {
		res = append(res, s[L])
	}
	return string(res)
}

func main() {
	S := "aabaccbc"
	r := reorganizeString(S)
	fmt.Println(r)
}
