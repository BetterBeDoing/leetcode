package leetcode

type LRUCache struct {
	head, tail *listNode
	cap        int
	mem        map[int]*listNode
}

type listNode struct {
	key, value int
	pre, nxt   *listNode
}

func newNode(key, value int) *listNode {
	return &listNode{
		key:   key,
		value: value,
		pre:   nil,
		nxt:   nil,
	}
}

func LRUConstructor(capacity int) LRUCache {
	tmp := new(LRUCache)
	tmp.cap = capacity
	tmp.head = newNode(0, 0)
	tmp.tail = newNode(0, 0)
	tmp.head.nxt = tmp.tail
	tmp.tail.pre = tmp.head
	tmp.mem = make(map[int]*listNode, capacity)
	return *tmp
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.mem[key]
	if !ok {
		return -1
	}
	this.reAddHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.mem[key]
	if !ok {
		node = newNode(key, value)
		this.addHead(node)
		if len(this.mem) > this.cap {
			this.removeTail()
		}
	} else {
		node.value = value
		this.reAddHead(node)
	}
}

func (this *LRUCache) addHead(node *listNode) {
	first := this.head.nxt

	this.mem[node.key] = node
	first.pre = node
	node.nxt = first

	this.head.nxt = node
	node.pre = this.head
}

func (this *LRUCache) removeTail() {
	preNode := this.tail.pre

	if preNode != this.head {
		this.removeNode(preNode)
		return
	}
}

func (this *LRUCache) removeNode(node *listNode) {
	delete(this.mem, node.key)
	preNode := node.pre
	nxtNode := node.nxt

	preNode.nxt = nxtNode
	nxtNode.pre = preNode
}

func (this *LRUCache) reAddHead(node *listNode) {
	this.removeNode(node)
	this.addHead(node)
}
