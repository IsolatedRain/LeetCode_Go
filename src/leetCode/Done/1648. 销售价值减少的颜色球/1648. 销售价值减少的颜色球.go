package main

import (
	"fmt"
	"sort"
)

func maxProfit(inventory []int, orders int) int {
	sort.Slice(inventory, func(i, j int) bool {
		return inventory[i] < inventory[j]
	})
	inventory = append(inventory, inventory[0]-1)
	balls := [][]int{}
	cnt := 1
	v := inventory[0]
	for i := 1; i < len(inventory); i++ {
		if inventory[i] == inventory[i-1] {
			cnt++
		} else {
			balls = append(balls, []int{cnt, v})
			cnt = 1
			v = inventory[i]
		}
	}
	res := 0
	mod := 1000000007
	for len(balls) > 0 {
		res %= mod
		curK := balls[len(balls)-1][0]
		curV := balls[len(balls)-1][1]
		nxV := 0
		balls = balls[:len(balls)-1]
		if len(balls) > 0 {
			nxV = balls[len(balls)-1][1]
		}
		if (curV-nxV)*curK <= orders {
			res += ((nxV + 1 + curV) * (curV - nxV) / 2) * curK
			if len(balls) > 0 {
				balls[len(balls)-1][0] += curK
			}
			orders -= (curV - nxV) * curK
		} else {
			a, b := orders/curK, orders%curK
			res += ((curV - a + 1 + curV) * a / 2) * curK
			res += (curV - a) * b
			return res % mod
		}
	}
	return res % mod
}

func main() {
	inventory := []int{2, 8, 4, 10, 6}
	// inventory := []int{2, 5}
	orders := 20
	r := maxProfit(inventory, orders)
	fmt.Println(r)
}
