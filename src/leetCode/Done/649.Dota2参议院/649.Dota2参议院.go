package main

import "fmt"

func predictPartyVictory(senate string) string {
	R := []int{}
	D := []int{}
	n := len(senate)
	for i, s := range senate {
		if s == 'D' {
			D = append(D, i)
		} else {
			R = append(R, i)
		}
	}
	for len(R) > 0 && len(D) > 0 {
		if R[0] < D[0] {
			R = append(R, R[0]+n)
		} else {
			D = append(D, D[0]+n)
		}
		R = R[1:]
		D = D[1:]
	}
	if len(R) > 0 {
		return "Radiant"
	}
	return "Dire"
}

func main() {
	senate := "RDD"
	r := predictPartyVictory(senate)
	fmt.Println(r)
}
