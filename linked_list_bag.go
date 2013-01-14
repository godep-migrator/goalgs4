package algs4

import (
	"reflect"
)

type LinkedListBag struct {
	first *Node
	n     int
}

type LinkedListBagCursor struct {
	n *Node
	b *LinkedListBag
}

func NewLinkedListBag() Bag {
	return &LinkedListBag{}
}

func (me *LinkedListBag) Add(item interface{}) {
	oldfirst := me.first
	me.first = &Node{
		item: item,
		next: oldfirst,
	}
	me.n++
}

func (me *LinkedListBag) IsEmpty() bool {
	return me.first == nil
}

func (me *LinkedListBag) Size() int {
	return me.n
}

func (me *LinkedListBag) First() IterableCursor {
	if me.IsEmpty() {
		return nil
	}

	return &LinkedListBagCursor{
		n: me.first,
		b: me,
	}
}

func (me *LinkedListBagCursor) Value() reflect.Value {
	return me.n.Value()
}

func (me *LinkedListBagCursor) Next() IterableCursor {
	next := me.n.next

	if next == nil {
		return nil
	}

	return &LinkedListBagCursor{
		n: next,
		b: me.b,
	}
}
