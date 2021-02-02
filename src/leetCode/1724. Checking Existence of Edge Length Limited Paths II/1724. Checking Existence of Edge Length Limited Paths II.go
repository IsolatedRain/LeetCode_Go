package main

// DistanceLimitedPathsExist 类
type DistanceLimitedPathsExist struct {
	nums []int
}

// Constructor 初始化
func Constructor(n int, edgeList [][]int) DistanceLimitedPathsExist {
	return DistanceLimitedPathsExist{[]int{}}
}

// Query 查询
func (d *DistanceLimitedPathsExist) Query(p int, q int, limit int) bool {
	return false
}

/**
 * Your DistanceLimitedPathsExist object will be instantiated and called as such:
 * obj := Constructor(n, edgeList);
 * param_1 := obj.Query(p,q,limit);
 */
