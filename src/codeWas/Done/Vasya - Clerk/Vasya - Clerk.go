// https://www.codewars.com/kata/555615a77ebc7c2c8a0000b8/train/go
package main

import "fmt"

// Tickets 能否找零
func Tickets(p []int) string {
	curChange := map[int]int{}
	for i := range p {
		if p[i] == 25 {
			curChange[25]++
		} else {
			if p[i] == 50 {
				if curChange[25] > 0 {
					curChange[25]--
					curChange[50]++
				} else {
					return "NO"
				}
			} else if p[i] == 100 {
				if curChange[50] > 0 {
					if curChange[25] > 0 {
						curChange[50]--
						curChange[25]--
					} else {
						return "NO"
					}
				} else {
					if curChange[25] >= 3 {
						curChange[25] -= 3
					} else {
						return "NO"
					}
				}
			}
		}
	}
	return "YES"
}

func main() {
	fmt.Println(Tickets([]int{25, 25, 50}))
}
