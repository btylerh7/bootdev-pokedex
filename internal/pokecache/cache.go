package pokecache

import "sync"

var (
	mu sync.Mutex
)

type Cache struct {
	entries map[string][]byte
}

func (c *Cache) Add(key string, val []byte) {
	mu.Lock()
	defer mu.Unlock()
	c.entries[key] = val

}
func (c *Cache) Get(key string) (val []byte, exists bool) {
	val, ok := c.entries[key]
	return val, ok

}

func NewCache() Cache {
	return Cache{
		entries: make(map[string][]byte),
	}

}
