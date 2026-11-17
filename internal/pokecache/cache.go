package pokecache

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu sync.Mutex
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
}

func (c *Cache) Add(key string, val []byte) {
	mu.Lock()
	defer mu.Unlock()
	entry := cacheEntry{
		time.Now(),
		val,
	}
	c.entries[key] = entry

}
func (c *Cache) Get(key string) (val []byte, exists bool) {
	mu.Lock()
	defer mu.Unlock()

	entry, ok := c.entries[key]
	return entry.val, ok

}

func (c *Cache) Reap() {
	mu.Lock()
	defer mu.Unlock()
	for key, entry := range c.entries {
		fmt.Println(entry.createdAt.Sub(time.Now()))
		if entry.createdAt.Sub(time.Now()) < 0 {
			delete(c.entries, key)
		}
	}

}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.Reap()
	}
}

func NewCache(interval time.Duration) Cache {
	entries := make(map[string]cacheEntry)
	cache := Cache{
		entries: entries,
	}
	go cache.reapLoop(interval)
	return cache

}
