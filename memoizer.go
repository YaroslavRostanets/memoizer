package cache

type Cache struct {
	memo map[string]any
}

func New() *Cache {
	return &Cache{
		memo: make(map[string]any),
	}
}

func (c *Cache) Get(key string) any {
	return c.memo[key]
}

func (c *Cache) Set(key string, value any) {
	c.memo[key] = value
}

func (c *Cache) Delete(key string) {
	delete(c.memo, key)
}
