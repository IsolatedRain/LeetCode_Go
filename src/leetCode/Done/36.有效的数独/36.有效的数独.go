package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	var checkRow func([]byte) bool
	var checkCol func([]byte) bool
	dot := byte('.')
	// 检查行是否有重复
	checkRow = func(curRow []byte) bool {
		curExisit := [9]bool{}
		for _, v := range curRow {
			if v == dot {
				continue
			}
			// fmt.Printf("%v, %v, %v\n", v, string(v), '0')
			idx := v - '0' - 1
			if curExisit[idx] {
				return false
			}
			curExisit[idx] = true
		}
		return true
	}

	checkCol = func(curCol []byte) bool {
		curExisit := [9]bool{}
		for _, v := range curCol {
			if v == dot {
				continue
			}
			idx := v - '0' - 1
			if curExisit[idx] {
				return false
			}
			curExisit[idx] = true
		}
		return true
	}

	// 检查小九宫
	pos := [3]int{0, 3, 6}
	var checkBox func() bool
	checkBox = func() bool {
		for _, i := range pos {
			for _, j := range pos {
				curExisit := [9]bool{}
				for r := i; r < i+3; r++ {
					for c := j; c < j+3; c++ {
						v := board[r][c]
						if v == dot {
							continue
						}
						idx := v - '0' - 1
						if curExisit[idx] {
							return false
						}
						curExisit[idx] = true
					}
				}
			}
		}
		return true
	}
	// 检查行/列
	for i := 0; i < 9; i++ {
		if !checkRow(board[i]) {
			return false
		}
		// 获取列
		curCol := []byte{}
		for j := 0; j < 9; j++ {
			curCol = append(curCol, board[j][i])
		}
		if !checkCol(curCol) {
			return false
		}
	}

	return checkBox()
}

func main() {
	// board := [][]string{
	// 	{"5", "3", ".", ".", "7", ".", ".", ".", "."},
	// 	{"6", ".", ".", "1", "9", "5", ".", ".", "."},
	// 	{".", "9", "8", ".", ".", ".", ".", "6", "."},
	// 	{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
	// 	{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
	// 	{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
	// 	{".", "6", ".", ".", ".", ".", "2", "8", "."},
	// 	{".", ".", ".", "4", "1", "9", ".", ".", "5"},
	// 	{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	// }
	board := [][]string{{".", ".", "4", ".", ".", ".", "6", "3", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {"5", ".", ".", ".", ".", ".", ".", "9", "."}, {".", ".", ".", "5", "6", ".", ".", ".", "."}, {"4", ".", "3", ".", ".", ".", ".", ".", "1"}, {".", ".", ".", "7", ".", ".", ".", ".", "."}, {".", ".", ".", "5", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}, {".", ".", ".", ".", ".", ".", ".", ".", "."}}
	bd := [][]byte{}
	fmt.Printf("输入:\n")
	for _, b := range board {
		fmt.Printf("%v\n", b)
		tmp := []byte{}
		for _, c := range b {
			tmp = append(tmp, []byte(c)...)
		}
		bd = append(bd, tmp)
	}
	r := isValidSudoku(bd)
	fmt.Println()
	fmt.Printf("输出: %v\n", r)
}
