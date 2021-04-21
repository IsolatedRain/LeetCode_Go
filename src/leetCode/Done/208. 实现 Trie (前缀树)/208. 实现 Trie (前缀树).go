package main

type Trie struct {
	tree  map[byte]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{tree: make(map[byte]*Trie)}
}
func newNode() *Trie {
	return &Trie{tree: make(map[byte]*Trie)}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	root := t
	for i := range word {
		if _, ok := root.tree[word[i]]; !ok {
			root.tree[word[i]] = newNode()
		}
		root = root.tree[word[i]]
	}
	root.isEnd = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	root := t
	for i := range word {
		if _, ok := root.tree[word[i]]; !ok {
			return false
		}
		root = root.tree[word[i]]
	}
	return root.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	root := t
	for i := range prefix {
		if _, ok := root.tree[prefix[i]]; !ok {
			return false
		}
		root = root.tree[prefix[i]]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
