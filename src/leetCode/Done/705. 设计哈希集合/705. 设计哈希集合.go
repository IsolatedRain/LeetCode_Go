// https://leetcode-cn.com/problems/design-hashset/
package main

import "container/list"

const base = 1000000007

// MyHashSet 哈希集合
type MyHashSet struct {
	data []list.List
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{make([]list.List, base)}

}
func (mhs *MyHashSet) hash(key int) int {
	return key % base

}

/** Returns true if this set contains the specified element */
func (mhs *MyHashSet) Contains(key int) bool {
	idx := mhs.hash(key)
	for item := mhs.data[idx].Front(); item != nil; item = item.Next() {
		if item.Value.(int) == key {
			return true
		}
	}
	return false
}

func (mhs *MyHashSet) Add(key int) {
	if !mhs.Contains(key) {
		idx := mhs.hash(key)
		mhs.data[idx].PushBack(key)
	}
}

func (mhs *MyHashSet) Remove(key int) {
	idx := mhs.hash(key)
	for item := mhs.data[idx].Front(); item != nil; item = item.Next() {
		if item.Value.(int) == key {
			mhs.data[idx].Remove(item)
		}
	}
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
