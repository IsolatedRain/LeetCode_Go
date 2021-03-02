package main

// NumMatrix 矩阵
type NumMatrix struct {
	d [][]int
}

// Constructor 始化矩阵
func Constructor(matrix [][]int) NumMatrix {
	row, col := len(matrix), len(matrix[0])
	dp := make([][]int, row+1)
	for r := range dp {
		dp[r] = make([]int, col+1)
	}
	for r := 1; r < row+1; r++ {
		for c := 1; c < col+1; c++ {
			dp[r][c] = dp[r-1][c] + dp[r][c-1] - dp[r-1][c-1] + matrix[r-1][c-1]
		}
	}
	return NumMatrix{dp}
}

// SumRegion 回矩阵区域和
func (mtx *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	dp := mtx.d
	return dp[row2+1][col2+1] - dp[row1][col2+1] - dp[row2+1][col1] + dp[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
