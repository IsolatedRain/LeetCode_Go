package main

import (
	"fmt"
	"sort"
	"strconv"
)

func largestMultipleOfThree(digits []int) string {
	sort.Slice(digits, func(i, j int) bool { return digits[i] > digits[j] })
	sum := 0
	for _, v := range digits {
		sum += v
	}
	remainder := sum % 3
	res := ""
	if remainder == 0 {
		for i := range digits {
			res += strconv.Itoa(digits[i])
		}
		return res
	}
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]%3 == remainder {
			for j := 0; j < i; j++ {
				res += strconv.Itoa(digits[j])
			}
			for j := i + 1; j < len(digits); j++ {
				res += strconv.Itoa(digits[j])
			}
			return res
		}
	}
	count := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]%3 == 3-remainder {
			digits[i] = -1
			count++
		}
		if count == 2 {
			break
		}
	}
	if count == 2 {
		for i := 0; i < len(digits); i++ {
			if digits[i] != -1 {
				res += strconv.Itoa(digits[i])
			}
		}
		return res
	}
	return res
}

func main() {
	// digits := []int{8, 6, 7, 1, 0}
	digits := []int{8, 1, 9}
	fmt.Println(largestMultipleOfThree(digits))
}
