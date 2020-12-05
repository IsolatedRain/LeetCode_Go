package main

import "fmt"

func confusingNumber(N int) bool {
	d := map[int]int{1: 1, 6: 9, 9: 6, 8: 8, 0: 0}
	newNum, curN := 0, N
	for curN > 0 {
		curNDigit := curN % 10
		v, ok := d[curNDigit]
		if !ok {
			return false
		}
		newNum = newNum*10 + v
		curN /= 10
	}
	return newNum != N
}

func main() {
	N := 88
	r := confusingNumber(N)
	fmt.Println(r)
}
