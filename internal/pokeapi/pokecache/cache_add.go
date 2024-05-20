package pokecache

import (
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	cacheEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mutex.Lock()
	c.items[key] = cacheEntry
	c.mutex.Unlock()
}
