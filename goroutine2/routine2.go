package main

import "fmt"

type Cache struct {
	set chan setKeyValue
	get chan string
	res chan string
	m   map[string]string
}

type setKeyValue struct {
	Key   string
	Value string
}

func (c *Cache) Initialize() {
	c.m = make(map[string]string)
	c.get = make(chan string)
	c.set = make(chan setKeyValue)
	c.res = make(chan string)

	go c.process()
}

func (c *Cache) process() {
	for {
		select {
		case msg := <-c.set:
			c.m[msg.Key] = msg.Value
		case msg := <-c.get:
			c.res <- c.m[msg]
		}
	}
}

func (c *Cache) Add(key, value string) {
	c.set <- setKeyValue{Key: key, Value: value}
}

func (c *Cache) Get(key string) string {
	c.get <- key
	res := <-c.res
	return res
}

func main() {
	var c Cache
	c.Initialize()

	c.Add("key1", "value1")
	c.Add("key2", "value2")
	fmt.Println(c.Get("key1"))
	fmt.Println(c.Get("key2"))
}
