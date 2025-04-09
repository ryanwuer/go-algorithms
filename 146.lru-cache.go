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

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end



