// https://www.codewars.com/kata/5c6d80d7ff00de2dcc4d4188/train/go
package main

import (
	"errors"
	"fmt"
	"sort"
)

// DiagonalsOfSquare 对角线输出
func DiagonalsOfSquare(input []string) ([]string, error) {
	res := []string{}
	err := errors.New("inValidInput")

	// valid check
	if len(input) == 0 {
		return res, err
	}
	for i := range input {
		if len(input[i]) != len(input) {
			return res, err
		}
		input[i] += fmt.Sprintf(" %d", i) // add index-info to tail
	}

	N := len(input)

	// sort by dictionary order.
	sort.Slice(input, func(i, j int) bool { return input[i] < input[j] })
	arr := append(input, input...)

	// get the required form  LeftTop -> rightBottom
	help := func(i int) string {
		ret := ""
		c := 0
		for x := i; x < i+N; x++ {
			ret += string(arr[x][c])
			c++
		}
		return arr[i][N+1:] + ret // add index-info to the head for sort
	}

	for i := range arr[:N] {
		res = append(res, help(i))
	}

	// sort by first chr , to original .
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	for i := range res {
		res[i] = res[i][1:]
	}

	return res, nil
}

func main() {
	fmt.Println(DiagonalsOfSquare([]string{"abcd", "kata", "1234", "qwer"}))
}
