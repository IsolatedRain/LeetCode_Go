package main

import "fmt"

func monotoneIncreasingDigits(N int) int {
	digits := []int{}
	curN := N
	for curN > 0 {
		cur := curN % 10
		curN /= 10
		digits = append(digits, cur)
	}
	n := len(digits)
	curMax := -1
	curID := 0
	for i := n - 1; i >= 0; i-- {
		if digits[i] > curMax {
			curMax = digits[i]
			curID = i
		}
		if i > 0 && digits[i] > digits[i-1] {
			break
		}
	}
	if curID == 0 {
		return N
	}
	for i := 0; i < curID; i++ {
		digits[i] = 9
	}
	digits[curID]--
	res := 0
	for i := n - 1; i >= 0; i-- {
		res = res*10 + digits[i]
	}
	return res
}

func main() {
	N := 12365
	r := monotoneIncreasingDigits(N)
	fmt.Println(r)
}
