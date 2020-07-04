package main

import "fmt"

type LRUCache struct {
	size, capacity int
	cache          map[int]*DoubleLinkedList
	head, tail     *DoubleLinkedList
}

type DoubleLinkedList struct {
	key, value int
	pre, next  *DoubleLinkedList
}

func initDoubleLinkedList(key, value int) *DoubleLinkedList {
	return &DoubleLinkedList{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    map[int]*DoubleLinkedList{},
		head:     initDoubleLinkedList(0, 0),
		tail:     initDoubleLinkedList(0, 0),
	}

	l.head.next = l.tail
	l.tail.pre = l.head

	return l
}

func (this *LRUCache) AddToHead(node *DoubleLinkedList) {
	node.pre = this.head
	node.next = this.head.next

	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DoubleLinkedList) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) moveToHead(node *DoubleLinkedList) {
	this.removeNode(node)
	this.AddToHead(node)
}

func (this *LRUCache) removeTail() *DoubleLinkedList {
	node := this.tail.pre

	this.removeNode(node)

	return node
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {

	if _, ok := this.cache[key]; ok {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	} else {
		node := initDoubleLinkedList(key, value)
		this.cache[key] = node
		this.AddToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4)) // 返回  4

}
