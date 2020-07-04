package main

type LRUCache struct {
	size, cap  int
	cache      map[int]*DoubleLinkedList
	head, tail *DoubleLinkedList
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
	lruCache := LRUCache{
		size:  0,
		cap:   capacity,
		head:  initDoubleLinkedList(0, 0),
		tail:  initDoubleLinkedList(0, 0),
		cache: map[int]*DoubleLinkedList{},
	}
	lruCache.head.next = lruCache.tail
	lruCache.tail.pre = lruCache.head

	return lruCache
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}

	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDoubleLinkedList(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.cap {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DoubleLinkedList) {
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
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DoubleLinkedList {
	node := this.tail.pre
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {

}
