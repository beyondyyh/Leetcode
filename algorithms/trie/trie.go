package trie

import _ "fmt"

// Trie是便于word插入和查找的数据结构
type Trie struct {
	// val  byte // 可以不需要
	sons [26]*Trie
	end  int
}

// NewTrie return *Trie
func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a'
		if node.sons[idx] == nil {
			// node.sons[idx] = &Trie{val: word[i]}
			node.sons[idx] = &Trie{}
		}
		// fmt.Printf("%+v %+v %+v\n", string(word[i]), idx, node.sons)
		node = node.sons[idx]
	}
	node.end++
}

func (t *Trie) Search(word string) bool {
	node := t
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a'
		if node.sons[idx] == nil {
			return false
		}
		node = node.sons[idx]
	}

	if node.end > 0 {
		return true
	}

	return false
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t
	size := len(prefix)
	for i := 0; i < size; i++ {
		idx := prefix[i] - 'a'
		if node.sons[idx] == nil {
			return false
		}
		node = node.sons[idx]
	}
	return true
}
