package main

import (
	"container/list"
	"fmt"
)

// LRU cache implementatoin with linked list for string/string

type LRUCache struct {
	capacity int
	cache    map[string]*list.Element
	list     *list.List
}

type cacheNode struct {
	key   string
	value string
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

func (lru *LRUCache) Get(key string) string {
	if elem, found := lru.cache[key]; found {
		lru.list.MoveToFront(elem)
		return elem.Value.(*cacheNode).value
	}
	return ""
}

func (lru *LRUCache) Put(key, value string) {
	if elem, found := lru.cache[key]; found {
		lru.list.MoveToFront(elem)
		elem.Value = value
	} else {
		if len(lru.cache) >= lru.capacity {
			// Remove the least recently used element from the cache
			oldest := lru.list.Back()
			if oldest != nil {
				delete(lru.cache, oldest.Value.(*cacheNode).key)
				lru.list.Remove(oldest)
			}
		}
		newNode := &cacheNode{key, value}
		newElem := lru.list.PushFront(newNode)
		lru.cache[key] = newElem
	}
}

func main() {
	cache := NewLRUCache(3)
	cache.Put("A", "AAA")
	cache.Put("B", "BBB")
	cache.Put("C", "CCC")

	fmt.Println(cache)
	cache.Put("D", "DDD")
	fmt.Println(cache)
	fmt.Println(cache.Get("C"))
	fmt.Println(cache)
}
