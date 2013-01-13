package algs4

import (
	"fmt"
)

type Node struct {
	item interface{}
	next *Node
}

func (me *Node) Value() interface{} {
	return me.item
}

func (me *Node) Int64() int64 {
	if v, ok := me.item.(int64); ok {
		return v
	}

	panic(fmt.Errorf("%v is not an int64!", me.item))
}

func (me *Node) Float64() float64 {
	if v, ok := me.item.(float64); ok {
		return v
	}

	panic(fmt.Errorf("%v is not a float64!", me.item))
}

func (me *Node) String() string {
	if v, ok := me.item.(string); ok {
		return v
	}

	return fmt.Sprintf("%v", me.item)
}

func (me *Node) Bool() bool {
	if v, ok := me.item.(bool); ok {
		return v
	}

	panic(fmt.Errorf("%v is not a bool!", me.item))
}

func (me *Node) Byte() byte {
	if v, ok := me.item.(byte); ok {
		return v
	}

	panic(fmt.Errorf("%v is not a byte!", me.item))
}

func (me *Node) Rune() rune {
	if v, ok := me.item.(rune); ok {
		return v
	}

	panic(fmt.Errorf("%v is not a rune!", me.item))
}
