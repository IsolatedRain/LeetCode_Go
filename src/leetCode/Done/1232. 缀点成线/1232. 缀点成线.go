package main

import "fmt"

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) == 2 {
		return true
	}

	x, y := coordinates[0][0], coordinates[0][1]
	deltaX, deltaY := coordinates[1][0]-x, coordinates[1][1]-y
	for _, v := range coordinates[2:] {
		curX, curY := v[0], v[1]
		if (curX-x)*deltaY != (curY-y)*deltaX {
			return false
		}
	}
	return true
}

func main() {
	coordinates := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}
	fmt.Println(checkStraightLine(coordinates))
}
