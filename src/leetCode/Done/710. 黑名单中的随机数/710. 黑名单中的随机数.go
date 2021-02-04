package main

import (
	"math/rand"
	"time"
)

// self.s = N - len(blacklist)
// # 小于s的黑名单元素集合
// b_lt_s = {i for i in blacklist if i < self.s}
// # 大于s的非黑名单元素集合
// # 等价于：w_gt_s = {i for i in range(self.s, N)} - set(blacklist)，感谢评论
// w_gt_s = {*range(self.s, N)} - set(blacklist)
// # 做映射，使用zip方便一点
// # 等价于：self.m = {k: v for k,v in zip(b_lt_s, w_gt_s)}
// self.m = dict(zip(b_lt_s, w_gt_s))
// def pick(self):
// """
// :rtype: int
// """
// r = randint(0, self.s-1)
// return r if r not in self.m else self.m[r]

// Solution 类
type Solution struct {
	size int
	d    map[int]int
}

// Constructor 初始化
func Constructor(N int, blacklist []int) Solution {
	bList := map[int]bool{}
	for i := range blacklist {
		bList[blacklist[i]] = true
	}
	b := []int{}
	for i := range blacklist {
		if blacklist[i] < N-len(blacklist) {
			b = append(b, blacklist[i])
		}
	}
	a := []int{}
	for i := N - len(blacklist); i < N; i++ {
		if !bList[i] {
			a = append(a, i)
		}
	}
	d := map[int]int{}
	for i := range a {
		d[b[i]] = a[i]
	}
	return Solution{N - len(blacklist), d}
}

// Pick 随机选择
func (s *Solution) Pick() int {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(s.size)
	if v, ok := s.d[i]; ok {
		return v
	}
	return int(i)
}
