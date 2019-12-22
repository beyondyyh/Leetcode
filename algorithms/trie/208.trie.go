package trie

import _ "fmt"

// Trie是便于word插入和查找的数据结构
type Trie struct {
	// val  byte // 可以不需要
	sons   [26]*Trie
	isLeaf bool
}

// NewTrie return *Trie
func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a' // get char code
		if node.sons[idx] == nil {
			// node.sons[idx] = &Trie{val: word[i]}
			node.sons[idx] = &Trie{}
		}
		// fmt.Printf("%+v %+v %+v\n", string(word[i]), idx, node.sons)
		node = node.sons[idx]
	}
	node.isLeaf = true
}

func (t *Trie) Search(word string) bool {
	node := t.findNodeWithWord(word)
	return node != nil && node.isLeaf
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.findNodeWithWord(prefix)
	return node != nil
}

func (t *Trie) findNodeWithWord(word string) *Trie {
	node := t
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a'
		if node.sons[idx] == nil {
			return nil
		}
		node = node.sons[idx]
	}
	return node
}
