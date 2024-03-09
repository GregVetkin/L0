package cache

import (
	"fmt"
	"l0/models/database"
	"log"
)

type Cache struct {
	cache	map[string] []byte
}


func (c *Cache) TakeOrder(uid string) ([]byte, error) {
	value, ok := c.cache[uid]
	if !ok {
		return nil, fmt.Errorf("No data with such key")
	}
	return value, nil
}

func (c *Cache) StoreOrder(uid string, data []byte) {
	c.cache[uid] = data
}

func (c *Cache) Fill() error {
	db, err := database.OpenDataBase()
	if err != nil {
		log.Print(err)
		return fmt.Errorf("Error")
	}
	
	orders, err := db.AllOrders()
	if err != nil {
		log.Print(err)
		return fmt.Errorf("Error")
	}
	
	for _, order := range orders {
		c.cache[order.Uid] = order.Data
	}
	
	return nil
}


func NewCache() Cache {
	var c Cache = Cache{cache: make(map[string] []byte)}
	return c
}

