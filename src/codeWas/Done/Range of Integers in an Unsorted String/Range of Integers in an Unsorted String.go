// https://www.codewars.com/kata/5b6b67a5ecd0979e5b00000e/train/go
package main

import (
	"fmt"
	"math"
	"strconv"
)

// MysteryRange 还原 打乱顺序后组合一起的连续数字
func MysteryRange(s string, n int) [2]int {
	sLen := len(s)
	nums := map[int]bool{}  // mark a num is exist or not
	curMax, curMin := math.MinInt16, math.MaxInt16  

	var dfs func(int) bool
	dfs = func(i int) bool {
		if i == sLen {
			// if reach to end, valid check:
			for k := curMin; k <= curMax; k++ {
				if !nums[k] { // if a num is not exist 
					return false
				}
			}
			return true
		}

		tmpMax, tmpMin := curMax, curMin  // for traceback
		for j := i + 1; j < i+4; j++ {
			if j > sLen { // out of s range
				break
			}

			if j-i > 1 && s[i] == '0' {  // 01, 02 is not a valid num
				break
			}

			curNum, _ := strconv.Atoi(s[i:j])
			if nums[curNum] {  // if already exist
				continue
			}

			if curNum > curMax {
				curMax = curNum
			}
			if curNum < curMin {
				curMin = curNum
			}

			if curMax-curMin >= n {  // if more than n
				curMax, curMin = tmpMax, tmpMin  // traceback
				continue
			}
			nums[curNum] = true  // mark
			if dfs(j) {
				return true
			}
			// traceback
			nums[curNum] = false
			curMax, curMin = tmpMax, tmpMin
		}
		return false
	}
	dfs(0)
	return [2]int{curMin, curMax}
}

func main() {
	// s := "1568141291110137"
	// n := 10
	// s := "10610211511099104113100116105103101111114107108112109"
	// n := 18
	// s := "21035768491"
	// n := 10
	s := "57755488799715021391892385890806723998334829144102629487228476633166132968378752953170191017641156224494559636997514311818440463689304227931374100687753602278209412179814735665358564258655"
	n := 94
	fmt.Println(MysteryRange(s, n))
}
