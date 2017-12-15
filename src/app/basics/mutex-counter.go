package main

import (
	"sync"
	"time"
	"fmt"
)

type SafeCounter struct {
	v map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	go func(){
		for i := 0; i < 10000000; i++ {
			go c.Inc("somekey")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c.Value("somekey"))
}
