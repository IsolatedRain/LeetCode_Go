package leetcode

func longestMountain(A []int) int {
	n := len(A)
	if n < 3 {
		return 0
	}
	left := make([]int, n)
	for i := 1; i < n; i++ {
		if A[i] > A[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	right := make([]int, n)
	for i := n - 2; i > 0; i-- {
		if A[i] > A[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	res := 0
	for i := 1; i < n-1; i++ {
		if left[i] > 0 && right[i] > 0 {
			res = max(res, left[i]+1+right[i])
		}
	}
	return res

}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
