package algs4

type ArrayQueue struct {
	items []interface{}
}

type ArrayQueueCursor struct {
	v interface{}
	q *ArrayQueue
	i int
}

func NewArrayQueue() Queue {
	return &ArrayQueue{
		items: []interface{}{},
	}
}

func (me *ArrayQueue) Enqueue(item interface{}) {
	me.items = append(me.items, item)
}

func (me *ArrayQueue) Dequeue() interface{} {
	if me.IsEmpty() {
		return nil
	}

	ret := me.items[0]
	me.items = me.items[1:]

	return ret
}

func (me *ArrayQueue) Size() int {
	return len(me.items)
}

func (me *ArrayQueue) IsEmpty() bool {
	return len(me.items) == 0
}

func (me *ArrayQueue) First() IterableCursor {
	if me.IsEmpty() {
		return nil
	}

	return &ArrayQueueCursor{
		v: me.items[0],
		i: 0,
		q: me,
	}
}

func (me *ArrayQueueCursor) Value() interface{} {
	return me.v
}

func (me *ArrayQueueCursor) Next() IterableCursor {
	i := me.i + 1

	if i >= len(me.q.items) {
		return nil
	}

	return &ArrayQueueCursor{
		v: me.q.items[i],
		i: i,
		q: me.q,
	}
}
