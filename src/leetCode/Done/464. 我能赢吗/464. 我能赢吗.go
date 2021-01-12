package main

import "fmt"

func canIWin(maxChoosableInteger int, desiredTotal int) bool {

	if desiredTotal < maxChoosableInteger {
		return true
	}
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}
	memo := make([]int, 1<<maxChoosableInteger)
	var dfs func(curMask, curDesireTotal int) bool
	dfs = func(curMask, curDesireTotal int) bool {
		if memo[curMask] != 0 {
			return memo[curMask] > 1
		}
		for i := 0; i < maxChoosableInteger; i++ {
			if curMask&(1<<i) == 0 {
				if curDesireTotal <= i+1 || !dfs(curMask|(1<<i), curDesireTotal-i-1) {
					memo[curMask] = 2
					return true
				}
			}
		}
		memo[curMask] = 1
		return false
	}
	return dfs(0, desiredTotal)
}

func main() {
	maxChoosableInteger := 5
	desiredTotal := 50
	fmt.Println(canIWin(maxChoosableInteger, desiredTotal))
}
