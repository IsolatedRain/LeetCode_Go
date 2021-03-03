// https://www.codewars.com/kata/566fc12495810954b1000030/train/go
package main

import "fmt"

// NbDig 计算 k*k 里 数字d的个数
func NbDig(n int, d int) int {
	res := 0
	var count func(num int) int
	count = func(num int) int {
		cnt := 0
		for num > 0 {
			m := num % 10
			if m == d {
				cnt++
			}
			num /= 10
		}
		return cnt
	}
	for i := 0; i <= n; i++ {
		res += count(i * i)
	}
	if d == 0 {
		res ++
	}
	return res
}

func main() {
	fmt.Println(NbDig(5750, 0) == 4700)
}
