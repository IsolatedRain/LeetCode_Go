//https://www.codewars.com/kata/555624b601231dc7a400017a/train/go
package main

import "fmt"

// JosephusSurvivor Josephus环
func JosephusSurvivor(n, k int) int {
	q := make([]int, n)
	for i := range q {
		q[i] = i + 1
	}
	count := 1
	for len(q) > 1 {
		curLen := len(q) - 1
		for i := 0; i < curLen; i++ {
			if count != k {
				q = append(q, q[i])
			} else {
				count = 0
			}
			count++
		}
		q = q[curLen:]
	}
	return q[0]
}

func main() {
	fmt.Println(JosephusSurvivor(2100000, 2))
}
