package main

// CombinationIterator 组合类
type CombinationIterator struct {
	id int
	d  []string
}

//Constructor 初始化
func Constructor(characters string, combinationLength int) CombinationIterator {
	n := len(characters)
	curBytes := make([]byte, combinationLength)
	res := []string{}
	var dfs func(i, k int)
	dfs = func(i, k int) {
		if k == combinationLength {
			res = append(res, string(curBytes))
			return
		}
		for j := i; j < n; j++ {
			curBytes[k] = characters[j]
			dfs(j+1, k+1)
		}
	}
	dfs(0, 0)
	return CombinationIterator{0, res}
}

// Next 返回 下一个组合
func (c *CombinationIterator) Next() string {
	res := c.d[c.id]
	c.id++
	return res
}

// HasNext if there's next comb.
func (c *CombinationIterator) HasNext() bool {
	return c.id != len(c.d)
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
