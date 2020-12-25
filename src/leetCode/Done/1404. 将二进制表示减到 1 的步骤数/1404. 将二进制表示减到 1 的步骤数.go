package main

import "fmt"

func numSteps(s string) int {
	x := []byte(s)
	count := 0
	for len(x) > 1 {
		if x[len(x)-1] == '1' {
			find := false
			zeroID := -1
			for i := len(x) - 1; i > -1; i-- {
				if x[i] == '0' {
					find = true
					zeroID = i
					break
				}
			}
			if !find {
				tmp := []byte{'1'}
				for range x {
					tmp = append(tmp, '0')
				}
				x = tmp
			} else {
				tmp := []byte{}
				for range x[zeroID+1:] {
					tmp = append(tmp, '0')
				}
				x[zeroID] = '1'
				x = append(x[:zeroID+1], tmp...)
			}
		} else {
			x = x[:len(x)-1]
		}
		count++
	}
	return count
}

func main() {
	s := "10110"
	fmt.Println(numSteps(s))
}
