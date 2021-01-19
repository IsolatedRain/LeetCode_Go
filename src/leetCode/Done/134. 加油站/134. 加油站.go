package main

import (
	"fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	remainGas := make([]int, n)
	for i := range gas {
		remainGas[i] = gas[i] - cost[i]
	}
	remainGas = append(remainGas, remainGas...)
	start := 0
	for start < n {
		if remainGas[start] < 0 {
			start++
			continue
		}
		curGas := 0
		circuit := true
		for i := start; i < start+n; i++ {
			curGas += remainGas[i]
			if curGas < 0 {
				circuit = false
				start = i + 1
				break
			}
		}
		if circuit {
			return start
		}
	}
	return -1
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}
