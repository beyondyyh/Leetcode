package linkedList

type LRUCache struct {
	capacity int
	hash     map[int]*DListNode
	Head     *DListNode
	Tail     *DListNode
}

func Constructor(capacity int) *LRUCache {
	head := &DListNode{Key: -1, Val: -1}
	tail := &DListNode{Key: -1, Val: -1}
	head.Next = tail
	tail.Prev = head

	cache := &LRUCache{
		capacity: capacity,
		hash:     make(map[int]*DListNode, capacity),
		Head:     head,
		Tail:     tail,
	}
	return cache
}

func (lru *LRUCache) insert(node *DListNode) {
	t := lru.Tail
	node.Prev = t.Prev
	t.Prev.Next = node
	node.Next = t
	t.Prev = node
}

func (lru *LRUCache) remove(node *DListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (lru *LRUCache) Len() int {
	return len(lru.hash)
}

func (lru *LRUCache) Cap() int {
	return lru.capacity
}

func (lru *LRUCache) Get(key int) int {
	if v, ok := lru.hash[key]; ok {
		lru.remove(v)
		lru.insert(v)
		return v.Val
	}
	return -1
}

func (lru *LRUCache) Put(key, val int) {
	if v, ok := lru.hash[key]; ok {
		lru.remove(v)
		lru.insert(v)
		v.Val = val
	} else {
		if lru.Len() == lru.Cap() {
			h := lru.Head.Next
			lru.remove(h)
			// Import! update lru.hash
			delete(lru.hash, h.Key)
		}
		node := &DListNode{
			Key: key,
			Val: val,
		}
		lru.hash[key] = node
		lru.insert(node)
	}
}
