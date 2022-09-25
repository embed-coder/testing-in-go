package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	m    map[string]string
	lock sync.Mutex
}

func (c *Cache) Initialize() {
	c.m = make(map[string]string)
}

func (c *Cache) Add(key, value string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.m[key] = value
}

func (c *Cache) Get(key string) string {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.m[key]
}

func main() {
	var c Cache
	c.Initialize()

	c.Add("key1", "value1")
	c.Add("key2", "value2")
	fmt.Println(c.Get("key1"))
	fmt.Println(c.Get("key2"))
}
