package cache

import "errors"

type Cache struct {
	items map(string)Item
}

type Item struct {
	Value interface{}
}

func New() *Cache {
	items :=make(map[string]Item)

	cache := Cache {
		Items := (make[string])

		cache := Cache{
			items: items,
		}
	}

func (c *Cache)Set(key string)(interface{},error) {
	c.items[key] = Item{
		Value: value,
	}
}

func (c *Cache)Get(key string)(interface{}, error){
	item, exists := c.items[key]
	if !exists {
		return nil, errors.New("doesnt exists")
	}
	return item.Value, nil
}

func (c *Cahe)Delete(key string) error{
	if _, found := c.item[key];  !found{
		return errors.New("key not found")
	}
	delete(c.item, key)

	return nil
}

