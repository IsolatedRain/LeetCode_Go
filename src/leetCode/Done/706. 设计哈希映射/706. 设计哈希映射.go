// https://leetcode-cn.com/problems/design-hashmap/
package main

import "container/list"

const base = 769

type dict struct {
	key, value int
}

type MyHashMap struct {
	data []list.List
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{make([]list.List, base)}
}
func (m MyHashMap) Hash(key int) int {
	return key % base
}

/** value will always be non-negative. */
func (m *MyHashMap) Put(key int, value int) {
	head := m.Hash(key)
	for node := m.data[head].Front(); node != nil; node = node.Next() {
		if d := node.Value.(dict); d.key == key {
			node.Value = dict{key, value}
			return
		}
	}
	m.data[head].PushBack(dict{key, value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (m *MyHashMap) Get(key int) int {
	head := m.Hash(key)
	for node := m.data[head].Front(); node != nil; node = node.Next() {
		if d := node.Value.(dict); d.key == key {
			return d.value
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (m *MyHashMap) Remove(key int) {
	head := m.Hash(key)
	for node := m.data[head].Front(); node != nil; node = node.Next() {
		if node.Value.(dict).key == key {
			m.data[head].Remove(node)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
