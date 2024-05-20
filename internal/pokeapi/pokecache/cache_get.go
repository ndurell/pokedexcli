package pokecache

func (c *Cache) Get(key string) (cacheEntry, bool) {
	c.mutex.RLock()
	v, ok := c.items[key]
	c.mutex.RUnlock()
	return v, ok
}
