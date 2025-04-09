/*
 * @lc app=leetcode.cn id=146 lang=golang
 * @lcpr version=30204
 *
 * [146] LRU 缓存
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
type LRUCache struct {
	// 这里可以使用 map 来存储 key 和对应的值
	// 也可以使用链表来存储 key 的顺序
	// 这里使用 map 来存储 key 和对应的值
	cache map[int]int
	// 这里使用一个双向链表来存储 key 的顺序
	// 这里使用一个切片来存储 key 的顺序
	keys []int
	// 容量
	capacity int
	// 当前缓存的大小
	size int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[int]int),
		keys:     make([]int, 0, capacity),
		capacity: capacity,
		size:     0,
	}
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.cache[key]; ok {
		// 如果 key 存在，则更新 key 的顺序
		for i, k := range this.keys {
			if k == key {
				this.keys = append(this.keys[:i], this.keys[i+1:]...)
				break
			}
		}
		this.keys = append(this.keys, key)
		return value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		// 如果 key 存在，则更新值
		this.cache[key] = value
		// 更新 key 的顺序
		for i, k := range this.keys {
			if k == key {
				this.keys = append(this.keys[:i], this.keys[i+1:]...)
				break
			}
		}
		this.keys = append(this.keys, key)
	} else {
		if this.size == this.capacity {
			// 如果缓存已满，则删除最久未使用的 key
			oldestKey := this.keys[0]
			delete(this.cache, oldestKey)
			this.keys = this.keys[1:]
			this.size--
		}
		this.cache[key] = value
		this.keys = append(this.keys, key)
		this.size++
	}
}

// Node 是双向链表节点
type Node struct {
    key, value int
    prev, next *Node
}

type LRUCache2 struct {
    capacity int
    cache    map[int]*Node
    head     *Node // dummy head
    tail     *Node // dummy tail
}

func Constructor2(capacity int) LRUCache2 {
    head := &Node{}
    tail := &Node{}
    head.next = tail
    tail.prev = head

    return LRUCache2{
        capacity: capacity,
        cache:    make(map[int]*Node),
        head:     head,
        tail:     tail,
    }
}

// 移除一个节点
func (this *LRUCache2) remove(node *Node) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

// 加到链表头
func (this *LRUCache2) insertToHead(node *Node) {
    node.next = this.head.next
    node.prev = this.head
    this.head.next.prev = node
    this.head.next = node
}

func (this *LRUCache2) Get2(key int) int {
    if node, ok := this.cache[key]; ok {
        // 移到头部
        this.remove(node)
        this.insertToHead(node)
        return node.value
    }
    return -1
}

func (this *LRUCache2) Put2(key int, value int) {
    if node, ok := this.cache[key]; ok {
        // 更新并移动到头部
        node.value = value
        this.remove(node)
        this.insertToHead(node)
    } else {
        if len(this.cache) == this.capacity {
            // 删除尾部最近最少使用节点
            lru := this.tail.prev
            this.remove(lru)
            delete(this.cache, lru.key)
        }
        newNode := &Node{key: key, value: value}
        this.cache[key] = newNode
        this.insertToHead(newNode)
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end



