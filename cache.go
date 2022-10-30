package cache

import "errors"

// Create struct with map
type Cache struct {
	items map[string]Item
}

// Create struct of value
type Item struct {
	Value interface{}
}

// Create constructor
func New() *Cache {
	items := make(map[string]Item)

	cache := Cache{
		items: items,
	}

	return &cache
}

// Add new item to cache
func (c *Cache) Set(key string, value interface{}) {
	c.items[key] = Item{
		Value: value,
	}
}

// Get item from cache storage
func (c *Cache) Get(key string) (interface{}, error) {
	item, exists := c.items[key]
	if !exists {
		return nil, errors.New("cache element not found")
	}
	return item.Value, nil
}

// Delete item from cache storage
func (c *Cache) Delete(key string) error {
	if _, found := c.items[key]; !found {
		return errors.New("key not found")
	}

	delete(c.items, key)

	return nil
}
