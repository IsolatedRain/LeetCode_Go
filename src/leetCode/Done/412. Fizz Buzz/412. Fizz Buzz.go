package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	res := make([]string, n+1)
	for i := 1; i < n+1; i++ {
		res[i] = strconv.Itoa(i)
	}
	for i := 2; i < n+1; i += 3 {
		res[i] = "Fizz"
	}
	for i := 4; i < n+1; i += 5 {
		if res[i] == "Fizz" {
			res[i] += "Buzz"
		} else {
			res[i] = "Buzz"
		}
	}
	return res[1:]
}

func main() {
	fmt.Println(fizzBuzz(15))
}
