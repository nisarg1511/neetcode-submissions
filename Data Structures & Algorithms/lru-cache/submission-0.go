type ListNode struct {
	Key   int
	Value int
	Prev  *ListNode
	Next  *ListNode
}

type LRUCache struct {
	Cache    map[int]*ListNode
	Head     *ListNode // Most recently used
	Elements int
	Capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cache:    make(map[int]*ListNode),
		Capacity: capacity,
	}
}

// Move node to head
func (this *LRUCache) moveToHead(node *ListNode) {
	if node == this.Head {
		return
	}

	// Remove node from current position
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	// Insert before current head
	tail := this.Head.Prev

	node.Next = this.Head
	node.Prev = tail

	tail.Next = node
	this.Head.Prev = node

	this.Head = node
}

func (this *LRUCache) Get(key int) int {
	node := this.Cache[key]

	if node == nil {
		return -1
	}

	this.moveToHead(node)

	return node.Value
}

func (this *LRUCache) Put(key int, value int) {

	// Key already exists
	if node := this.Cache[key]; node != nil {
		node.Value = value
		this.moveToHead(node)
		return
	}

	// First insertion
	if this.Elements == 0 {
		node := &ListNode{
			Key:   key,
			Value: value,
		}

		node.Next = node
		node.Prev = node

		this.Head = node
		this.Cache[key] = node
		this.Elements++

		return
	}

	// Cache full -> remove LRU (tail)
	if this.Elements == this.Capacity {
		tail := this.Head.Prev

		// Single node case
		if this.Elements == 1 {
			delete(this.Cache, tail.Key)

			tail.Key = key
			tail.Value = value

			this.Cache[key] = tail
			this.Head = tail

			return
		}

		// Remove tail
		tail.Prev.Next = this.Head
		this.Head.Prev = tail.Prev

		delete(this.Cache, tail.Key)

		// Reuse tail node as new head
		tail.Key = key
		tail.Value = value

		tail.Next = this.Head
		tail.Prev = this.Head.Prev

		this.Head.Prev.Next = tail
		this.Head.Prev = tail

		this.Head = tail

		this.Cache[key] = tail

		return
	}

	// Insert new node at head
	tail := this.Head.Prev

	node := &ListNode{
		Key:   key,
		Value: value,
		Next:  this.Head,
		Prev:  tail,
	}

	tail.Next = node
	this.Head.Prev = node

	this.Head = node

	this.Cache[key] = node
	this.Elements++
}