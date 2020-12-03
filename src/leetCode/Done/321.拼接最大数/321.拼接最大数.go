package main

import "fmt"

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	res := []int{}
	n1 := len(nums1)
	end := k
	if n1 < k {
		end = n1
	}
	for i := 0; i <= end; i++ {
		s1 := getMaxSubSeq(nums1, i)
		s2 := getMaxSubSeq(nums2, k-i)
		fmt.Println(s1, s2, i)
		curSeq := merge(s1, s2)
		if len(curSeq) < k {
			continue
		}
		if compareSeqs(curSeq, res) {
			res = curSeq
		}
	}
	return res
}
func getMaxSubSeq(seq []int, k int) []int {
	stack := []int{}
	for i, v := range seq {
		for len(stack) > 0 && stack[len(stack)-1] < v && len(stack)+len(seq)-1-i >= k {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
	if k > len(stack) {
		return stack
	}
	return stack[:k]
}

// 判断 seq1 > seq2
func compareSeqs(seq1, seq2 []int) bool {
	for i := 0; i < len(seq1) && i < len(seq2); i++ {
		if seq1[i] != seq2[i] {
			return seq1[i] > seq2[i]
		}
	}
	return len(seq1) > len(seq2)
}

// 合并返回 seq1 seq2组合的最大序列
func merge(seq1, seq2 []int) []int {
	n1 := len(seq1)
	n2 := len(seq2)
	res := make([]int, n1+n2)
	id1 := 0
	id2 := 0
	i := 0
	for id1 < n1 && id2 < n2 {
		if seq1[id1] > seq2[id2] {
			res[i] = seq1[id1]
			id1++
		} else if seq1[id1] < seq2[id2] {
			res[i] = seq2[id2]
			id2++
		} else {
			if compareSeqs(seq1[id1:], seq2[id2:]) {
				res[i] = seq1[id1]
				id1++
			} else {
				res[i] = seq2[id2]
				id2++
			}
		}
		i++
	}
	if id1 != n1 {
		res = append(res[:i], seq1[id1:]...)
	} else {
		res = append(res[:i], seq2[id2:]...)
	}
	return res
}

func main() {
	// nums1 := []int{3, 4, 6, 5}
	// nums2 := []int{9, 1, 2, 5, 8, 3}
	// k := 2
	// nums1 := []int{3, 3, 1, 8, 2, 4, 2, 9, 2, 4, 7, 1, 9, 2, 3, 4, 0, 7, 5, 4}
	// nums2 := []int{9, 7, 7, 1, 3, 6, 8, 6, 9, 6, 0, 4, 3, 6, 6, 1, 0, 4, 6, 2, 2, 6, 4, 6, 0, 4, 9, 7, 4, 9, 8, 4, 9, 8, 4, 6, 6, 5, 8, 2, 8, 6, 6, 6, 1, 0, 9, 0, 8, 0, 4, 0, 4, 4, 1, 7, 9, 8, 4, 2, 2, 0, 3, 2, 3, 9, 1, 8, 9, 5, 2, 7, 9, 2, 7, 7, 8, 5, 4, 4, 8, 6, 5, 5, 9, 6, 1, 4, 6, 0, 8, 5, 3, 4, 2, 0, 0, 9, 5, 2}
	// k := 100
	nums1 := []int{3, 9}
	nums2 := []int{8, 9}
	k := 3
	r := maxNumber(nums1, nums2, k)
	fmt.Printf("输出：%v\n", r)
}
