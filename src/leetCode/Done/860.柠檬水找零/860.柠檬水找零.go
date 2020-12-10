package main

import "fmt"

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, b := range bills {
		if b == 5 {
			five++
		} else if b == 10 {
			five--
		} else {
			if ten > 0 {
				ten--
				five--
			} else {
				five -= 3
			}
		}
		if five < 0 {
			return false
		}
	}
	return true
}

func main() {
	bills := []int{5, 5, 5, 10, 20, 20}
	r := lemonadeChange(bills)
	fmt.Println(r)
}
