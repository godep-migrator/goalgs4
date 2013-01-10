package algs4

import (
	"fmt"
)

type Counter struct {
	name  string
	count int
}

func NewCounter(name string) *Counter {
	counter := &Counter{}
	counter.name = name
	return counter
}

func (c *Counter) Increment() {
	c.count += 1
}

func (c *Counter) Tally() int {
	return c.count
}

func (c *Counter) String() string {
	return fmt.Sprintf("%d %s", c.count, c.name)
}

func (c *Counter) Cmp(other *Counter) int {
	if c.count < other.count {
		return -1
	} else if c.count > other.count {
		return 1
	}
	return 0
}
