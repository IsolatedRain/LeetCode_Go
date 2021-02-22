package main

import "fmt"

// diff[i] 表示到 差分数组， L-R 翻转1次， 则在R+1 处再翻转一次 消除影响。
// pre 表示累计翻转了几次, 偶数=没翻转， 奇数=翻转1次
func minKBitFlips(A []int, K int) int {
	n, res, pre := len(A), 0, 0
	diff := make([]int, n+1)
	for i, v := range A {
		pre += diff[i]
		pre %= 2
		if pre^v == 0 {
			if i+K > n {
				return -1
			}
			res++
			diff[i+K]++
			pre++
		}
	}
	return res
}

func main() {
	A := []int{0, 0, 0, 1, 0, 1, 1, 0}
	K := 3
	fmt.Println(minKBitFlips(A, K))
}
