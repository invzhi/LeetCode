// 146. LRU Cache

// Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.
// get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
// put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

// Follow up:
// Could you do both operations in O(1) time complexity?

// Example:
// LRUCache cache = new LRUCache( 2 /* capacity */ );
// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // returns 1
// cache.put(3, 3);    // evicts key 2
// cache.get(2);       // returns -1 (not found)
// cache.put(4, 4);    // evicts key 1
// cache.get(1);       // returns -1 (not found)
// cache.get(3);       // returns 3
// cache.get(4);       // returns 4

package leetcode

import "container/list"

type keyValue struct {
	key, value int
}

// LRUCache implement Least Recently Used (LRU) cache.
type LRUCache struct {
	l *list.List
	m map[int]*list.Element
	c int
}

// Constructor creates and initializes a new LRUCache using capacity as its capacity.
func Constructor(capacity int) LRUCache {
	return LRUCache{
		l: list.New(),
		m: make(map[int]*list.Element),
		c: capacity,
	}
}

// Get get the value of the key if the key exists in the cache, otherwise return -1.
func (c *LRUCache) Get(key int) int {
	e, ok := c.m[key]
	if !ok {
		return -1
	}
	c.l.MoveToFront(e)
	return e.Value.(*keyValue).value
}

// Put set or insert the value if the key is not already present.
func (c *LRUCache) Put(key, value int) {
	e, ok := c.m[key]
	if ok {
		c.l.MoveToFront(e)
		e.Value.(*keyValue).value = value
	} else {
		e = c.l.PushFront(&keyValue{key, value})
		c.m[key] = e
		if c.l.Len() > c.c {
			oldestElem := c.l.Back()
			c.l.Remove(oldestElem)
			delete(c.m, oldestElem.Value.(*keyValue).key)
		}
	}
}
