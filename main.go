package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	A int
	B int
}

func (c *Counter) increment() {
	c.A++
	c.B++
}

func NewCounter(a int, b int) *Counter {
	c := new(Counter)
	c.A = a
	c.B = b
	return c
}

var counterPool = sync.Pool{
	New: func() interface{} { return new(Counter) },
}

func main() {
	cc := NewCounter(2, 3)
	fmt.Println("init: ", cc)
	cc.increment()
	fmt.Println("final: ", cc)
}
