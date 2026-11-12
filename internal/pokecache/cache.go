package pokecache

import (
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

}

func (c *Cache) reapLoop(interval time.Duration) {
	// ticker := time.NewTicker(interval)
}

func NewCache(interval time.Duration) Cache {
	var entries map[string]cacheEntry
	return Cache{
		entries: entries,
	}

}
