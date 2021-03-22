//https://leetcode-cn.com/problems/flatten-nested-list-iterator/

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
package main

type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool { return true }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int { return 0 }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger { return []*NestedInteger{} }

type NestedIterator struct {
	data []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var dfs func(curNest []*NestedInteger) []int
	dfs = func(curNest []*NestedInteger) []int {
		res := []int{}
		for _, nl := range curNest {
			if nl.IsInteger() {
				res = append(res, nl.GetInteger())
			} else {
				res = append(res, dfs(nl.GetList())...)
			}
		}
		return res
	}
	nums := dfs(nestedList)
	return &NestedIterator{data: nums}
}

func (nIt *NestedIterator) Next() int {
	val := nIt.data[0]
	nIt.data = nIt.data[1:]
	return val
}

func (nIt *NestedIterator) HasNext() bool {
	return len(nIt.data) > 0
}
