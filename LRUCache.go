package main

import (
	"fmt"
)

type Node struct {
	value, key int
	next, prev *Node
}

type LRUCache struct {
	cache map[int]*Node
	head, tail *Node
	capacity int
}

func Constructor(capacity int) LRUCache {
	p := new(LRUCache)
	p.capacity = capacity
	p.cache = map[int]*Node{}
	
	return *p
}

func (LRU *LRUCache) Get(key int) int {
	if node, ok := LRU.cache[key]; ok {
		LRU.Evict(node)
		LRU.Insert(node)
		return node.value
	}

	return -1
}

func (LRU *LRUCache) Put(key int, value int) {
	if node, ok := LRU.cache[key]; ok {
		node.value = value
		LRU.Evict(node)
		LRU.Insert(node)
	} else {
		node := &Node{value: value, key: key}
		LRU.Insert(node)
		LRU.cache[key] = node
	}

	if len(LRU.cache) > LRU.capacity {
		lru := LRU.head
		LRU.Evict(lru)
		delete(LRU.cache, lru.key)
	}

}

func (LRU *LRUCache) Evict(node *Node) {
	if node.next != nil {
		node.next.prev = node.prev
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if LRU.head == node {
		LRU.head = node.next
	}

	if LRU.tail == node {
		LRU.tail = node.prev
	}

	node.next = nil
	node.prev = nil
}

func (LRU *LRUCache) Insert(node *Node) {
	if LRU.tail == nil {
		LRU.tail = node
		LRU.head = LRU.tail
		return
	}

	LRU.tail.next = node
	node.prev = LRU.tail
	LRU.tail = node
}

func main() {
	fmt.Println("yooo")
}