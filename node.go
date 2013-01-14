package algs4

import (
	"reflect"
)

type Node struct {
	item interface{}
	next *Node
}

func (me *Node) Value() reflect.Value {
	return reflect.ValueOf(me.item)
}
