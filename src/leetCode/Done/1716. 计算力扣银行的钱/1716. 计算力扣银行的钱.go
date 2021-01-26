package main

import "fmt"

func totalMoney(n int) int {
	if n < 8 {
		return (1 + n) * n / 2
	}
	oneWeek := (1 + 7) * 7 / 2
	weeks, days := n/7, n%7
	return (oneWeek+(oneWeek+7*(weeks-1)))*(weeks)/2 + (weeks+1+weeks+1+days-1)*days/2
}

func main() {
	fmt.Println(totalMoney(20))
}
