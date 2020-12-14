package main

import (
	"fmt"
	"sort"
)

//V 存价值总和， Alice Bob对应的价值
type V struct {
	valSum int
	aVal   int
	bVal   int
}

// 对手减少的分 等价于 自己增加的分
func stoneGameVI(aliceValues []int, bobValues []int) int {
	n := len(aliceValues)
	values := make([]V, n, n)
	for i := range aliceValues {
		values[i].valSum = aliceValues[i] + bobValues[i]
		values[i].aVal = aliceValues[i]
		values[i].bVal = bobValues[i]
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i].valSum > values[j].valSum
	})
	aV, bV := 0, 0
	for i := range values {
		if i&1 == 1 {
			bV += values[i].bVal
		} else {
			aV += values[i].aVal
		}
	}
	if aV > bV {
		return 1
	} else if aV == bV {
		return 0
	}
	return -1
}

func main() {
	aliceValues := []int{1, 2}
	bobValues := []int{3, 1}
	r := stoneGameVI(aliceValues, bobValues)
	fmt.Println(r)

}
