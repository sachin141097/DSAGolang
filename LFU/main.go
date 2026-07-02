package main

import "fmt"

type node struct {
	key, value, freq int
	prev, next       *node
}

type dll struct {
	head, tail *node
	size       int
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
	d.size++
}

func (d *dll) deleteNode(n *node) {
	prev, next := n.prev, n.next
	prev.next = next
	next.prev = prev
	d.size--
}

func (d *dll) removeLast() *node {
	if d.size == 0 {
		return nil
	}
	n := d.tail.prev
	d.deleteNode(n)
	return n
}

type LFUCache struct {
	keyMap   map[int]*node
	freqMap  map[int]*dll
	capacity int
	minFreq  int
}

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		keyMap:   make(map[int]*node),
		freqMap:  make(map[int]*dll),
		capacity: capacity,
	}
}
func (c *LFUCache) updateFrequency(n *node) {
	oldFreq := n.freq
	oldList := c.freqMap[oldFreq]
	oldList.deleteNode(n)
	if oldList.size == 0 {
		delete(c.freqMap, oldFreq)
		if oldFreq == c.minFreq {
			c.minFreq++
		}
	}
	n.freq++
	newList, ok := c.freqMap[n.freq]
	if !ok {
		newList = newDLL()
		c.freqMap[n.freq] = newList
	}
	newList.addNode(n)
}
func (c *LFUCache) Get(key int) (int, bool) {
	n, ok := c.keyMap[key]
	if !ok {
		return 0, false
	}
	c.updateFrequency(n)
	return n.value, true
}
func (c *LFUCache) Put(key, value int) {
	if c.capacity == 0 {
		return
	}
	if n, ok := c.keyMap[key]; ok {
		n.value = value
		c.updateFrequency(n)
		return
	}
	if len(c.keyMap) == c.capacity {
		minList := c.freqMap[c.minFreq]
		victim := minList.removeLast()
		delete(c.keyMap, victim.key)
		if minList.size == 0 {
			delete(c.freqMap, c.minFreq)
		}
	}
	newNode := &node{key: key, value: value, freq: 1}
	list, ok := c.freqMap[1]
	if !ok {
		list = newDLL()
		c.freqMap[1] = list
	}
	list.addNode(newNode)
	c.keyMap[key] = newNode
	c.minFreq = 1
}
func (c *LFUCache) PrintState() {
	fmt.Println("--------------------------------")

	for freq, list := range c.freqMap {
		fmt.Printf("Freq %d : ", freq)
		curr := list.head.next
		for curr != list.tail {
			fmt.Printf("(%d,%d) ", curr.key, curr.value)
			curr = curr.next
		}
		fmt.Println()
	}

	fmt.Println("minFreq =", c.minFreq)
	fmt.Println("--------------------------------")
}
func main() {
	cache := NewLFUCache(2)
	cache.Put(1, 10)
	cache.Put(1, 20)
	cache.PrintState()
	cache.Put(2, 20)
	cache.PrintState()
	cache.Put(3, 30)
	cache.PrintState()
}
