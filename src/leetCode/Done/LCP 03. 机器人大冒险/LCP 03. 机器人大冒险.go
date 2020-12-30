package main

import (
	"fmt"
	"sort"
)

func robot(command string, obstacles [][]int, x int, y int) bool {
	sort.Slice(obstacles, func(i, j int) bool {
		return obstacles[i][0] < obstacles[j][0]
	})
	move := map[byte][]int{'U': {0, 1}, 'R': {1, 0}}
	// 一轮命令 能到的坐标
	d := map[[2]int]struct{}{}
	maxX, maxY := 0, 0
	curX, curY := 0, 0
	d[[2]int{curX, curY}] = struct{}{}
	for i := range command {
		m := move[command[i]]
		nxX := curX + m[0]
		nxY := curY + m[1]
		d[[2]int{nxX, nxY}] = struct{}{}
		curX, curY = nxX, nxY
		maxX += m[0]
		maxY += m[1]
	}
	
	// 检查是否在路径上，
	// 等比缩放
	var check func(x, y int) bool
	check = func(x, y int) bool {
		loops := min(x/maxX, y/maxY)
		newX, newY := x-(loops*maxX), y-(loops*maxY)
		key := [2]int{newX, newY}
		if _, ok := d[key]; ok {
			return false
		}
		return true
	}
	for i := range obstacles {
		if obstacles[i][0] > x || obstacles[i][1] > y {
			break
		}
		if !check(obstacles[i][0], obstacles[i][1]) {
			return false
		}
	}
	return !check(x, y)
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	command := "RRRUUU"
	obstacles := [][]int{{3, 0}, {3, 1}}
	x := 3
	y := 3
	r := robot(command, obstacles, x, y)
	fmt.Println(r)
}
