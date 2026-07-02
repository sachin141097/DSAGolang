package main

import "fmt"

type node struct {
	key, value int
	prev, next *node
}

type dll struct {
	head, tail *node
}

func newDLL() *dll {
	head := &node{key: -1, value: -1}
	tail := &node{key: -1, value: -1}
	head.next = tail
	tail.prev = head
	return &dll{head: head, tail: tail}
}
func (d *dll) addNode(n *node) {
	temp := d.head.next
	n.next = temp
	n.prev = d.head
	d.head.next = n
	temp.prev = n

}
func (d *dll) deleteNode(n *node) {
	prev := n.prev
	next := n.next
	prev.next = next
	next.prev = prev

}

type LRUCache struct {
	keyMap   map[int]*node
	capacity int
	list     *dll
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		keyMap:   make(map[int]*node),
		capacity: capacity,
		list:     newDLL(),
	}
}
func (c *LRUCache) moveToFront(n *node) {
	c.list.deleteNode(n)
	c.list.addNode(n)
}
func (c *LRUCache) Put(key, value int) {
	if c.capacity == 0 {
		return
	}
	if n, ok := c.keyMap[key]; ok {
		n.value = value
		c.moveToFront(n)
		return
	}
	if len(c.keyMap) == c.capacity {
		lru := c.list.tail.prev
		c.list.deleteNode(lru)
		delete(c.keyMap, lru.key)
	}
	n := &node{key: key, value: value}
	c.list.addNode(n)
	c.keyMap[key] = n

}
func (c *LRUCache) Get(key int) (int, bool) {
	n, ok := c.keyMap[key]
	if !ok {
		return 0, false
	}
	c.moveToFront(n)
	return n.value, true

}
func (c *LRUCache) printCache() {
	fmt.Println("--------------------------------")
	curr := c.list.head.next
	for curr != c.list.tail {
		fmt.Printf("(%d,%d) ", curr.key, curr.value)
		curr = curr.next
	}
	fmt.Println()
	fmt.Println("--------------------------------")

}
func main() {
	cache := NewLRUCache(4)
	cache.Put(1, 10)
	cache.Put(2, 20)
	cache.Put(3, 30)
	cache.Put(4, 40)
	cache.printCache()
	cache.Get(1)
	cache.printCache()
	cache.Put(5, 50)
	cache.printCache()

}
