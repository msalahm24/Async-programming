package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu      sync.Mutex
	counter map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter[name]++
}

func main() {
	c := Container{
		mu:      sync.Mutex{},
		counter: map[string]int{},
	}

	var wg sync.WaitGroup

	doIncre := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}
	wg.Add(4)
	go doIncre("a", 100)
	go doIncre("b", 100)
	go doIncre("a", 100)
	go doIncre("b", 100)
	wg.Wait()

	fmt.Println(c.counter)
}
