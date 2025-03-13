package cache

type Cache struct {
	cache map[string]interface{}
}

func New() *Cache {
	return &Cache{make(map[string]interface{})}
}

func (c *Cache) Set(key string, value interface{}) {
	c.cache[key] = value
}

func (c *Cache) Get(key string) interface{} {
	var value, ok = c.cache[key]

	if !ok {
		return "unknown key"
	}

	return value
}

func (c *Cache) Delete(key string) {
	delete(c.cache, key)
}
