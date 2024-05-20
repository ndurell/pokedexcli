package pokecache

import (
	"fmt"
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
	defer c.reapLoop()
	return c
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	go func() {
		for t := range ticker.C {
			c.mutex.Lock()
			for k, v := range c.items {
				fmt.Println(t.Sub(v.createdAt))
				if t.Sub(v.createdAt) > c.interval {
					delete(c.items, k)
				}
			}
			c.mutex.Unlock()
		}
	}()
}
