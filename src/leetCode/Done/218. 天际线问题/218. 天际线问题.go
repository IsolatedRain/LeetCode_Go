package main

import "fmt"

func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	if n == 1 {
		L, R, y := buildings[0][0], buildings[0][1], buildings[0][2]
		return [][]int{{L, y}, {R, 0}}
	}
	left := getSkyline(buildings[:n/2])
	right := getSkyline(buildings[n/2:])
	return merge(left, right)
}
func merge(left, right [][]int) (output [][]int) {
	var update func(x, y int)
	update = func(x, y int) {
		n := len(output)
		if n == 0 || output[n-1][0] != x {
			output = append(output, []int{x, y})
		} else {
			output[n-1][1] = y
		}
	}
	var updateRest func(curID, curY int, curSkyLine [][]int)
	updateRest = func(curID, curY int, curSkyLine [][]int) {
		for curID < len(curSkyLine) {
			x, y := curSkyLine[curID][0], curSkyLine[curID][1]
			if curY != y {
				update(x, y)
				curY = y
			}
			curID++
		}
	}
	leftID, rightID, leftLen, rightLen := 0, 0, len(left), len(right)
	leftY, rightY, curY, x := 0, 0, 0, 0
	for leftID < leftLen && rightID < rightLen {
		leftPoint, rightPoint := left[leftID], right[rightID]
		if leftPoint[0] < rightPoint[0] {
			x, leftY = leftPoint[0], leftPoint[1]
			leftID++
		} else {
			x, rightY = rightPoint[0], rightPoint[1]
			rightID++
		}
		maxY := max(leftY, rightY)
		if maxY != curY {
			update(x, maxY)
			curY = maxY
		}
	}
	if leftID == len(left) {
		updateRest(rightID, curY, right)
	} else {
		updateRest(leftID, curY, left)
	}
	return
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y

}

func main() {
	buildings := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	fmt.Println(getSkyline(buildings))
}
