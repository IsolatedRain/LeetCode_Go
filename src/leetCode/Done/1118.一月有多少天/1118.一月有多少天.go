package main

import "fmt"

func numberOfDays(Y int, M int) int {
	// 百年不润，四百年一润，四年一润。
	leapYear := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	notLeapYear := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if (Y%100 != 0 && Y%4 == 0) || (Y%400 == 0) {
		return leapYear[M-1]
	}
	return notLeapYear[M-1]
}

func main() {
	Y := 1992
	M := 7
	r := numberOfDays(Y, M)
	fmt.Println(r)
}
