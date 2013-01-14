package algs4

import (
	"reflect"
)

type LinkedListQueue struct {
	first *Node
	last  *Node
	n     int
}

type LinkedListQueueCursor struct {
	n *Node
	q *LinkedListQueue
}

func NewLinkedListQueue() Queue {
	return &LinkedListQueue{}
}

func (me *LinkedListQueue) Enqueue(item interface{}) {
	oldlast := me.last
	me.last = &Node{item: item}

	if me.IsEmpty() {
		me.first = me.last
	} else {
		oldlast.next = me.last
	}

	me.n++
}

func (me *LinkedListQueue) Dequeue() interface{} {
	if me.IsEmpty() {
		return nil
	}

	item := me.first.item
	oldfirst := me.first
	me.first = oldfirst.next
	oldfirst.next = nil
	me.n--

	if me.IsEmpty() {
		me.last = nil
	}

	return item
}

func (me *LinkedListQueue) Size() int {
	return me.n
}

func (me *LinkedListQueue) IsEmpty() bool {
	return me.n == 0
}

func (me *LinkedListQueue) First() IterableCursor {
	if me.IsEmpty() {
		return nil
	}

	return &LinkedListQueueCursor{
		n: me.first,
		q: me,
	}
}

func (me *LinkedListQueueCursor) Value() reflect.Value {
	return me.n.Value()
}

func (me *LinkedListQueueCursor) Next() IterableCursor {
	next := me.n.next

	if next == nil {
		return nil
	}

	return &LinkedListQueueCursor{
		n: next,
		q: me.q,
	}
}
