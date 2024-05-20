package pokecache

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	v, ok := c.items[key]
	c.mutex.RUnlock()
	return v.val, ok
}
