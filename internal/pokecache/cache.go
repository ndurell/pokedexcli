package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	items    map[string]cacheEntry
	interval time.Duration
	mutex    *sync.RWMutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		items:    make(map[string]cacheEntry),
		interval: interval,
		mutex:    &sync.RWMutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mutex.Lock()
	for k, v := range c.items {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.items, k)
		}
	}
	c.mutex.Unlock()
}
