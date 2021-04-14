package main

type Node struct {
	isWord bool
	next   map[byte]*Node
}

type Trie struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor3() Trie {
	return Trie{root: &Node{
		isWord: false,
		next:   make(map[byte]*Node),
	}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this.root

	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := cur.next[c]; !ok {
			node := &Node{
				next: make(map[byte]*Node),
			}
			cur.next[c] = node
		}
		cur = cur.next[c]
	}

	cur.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this.root

	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := cur.next[c]; !ok {
			return false
		}
		cur = cur.next[c]
	}

	return cur.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root

	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		if _, ok := cur.next[c]; !ok {
			return false
		}
		cur = cur.next[c]
	}

	return true
}
