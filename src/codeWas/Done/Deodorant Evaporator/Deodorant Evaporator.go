// https://www.codewars.com/kata/5506b230a11c0aeab3000c1f/go

package main

import "fmt"

// Evaporator 浓度降低
func Evaporator(content float64, evapPerDay int, threshold int) int {
	res := 0
	curEvap := float64(100-evapPerDay) / 100
	decPerDay := curEvap
	t := float64(threshold) / 100
	for curEvap >= t {
		curEvap *= decPerDay
		res++
	}
	return res + 1
}

func main() {
	fmt.Println(Evaporator(10, 10, 5), 29)
}
