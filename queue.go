package algs4

type Queue struct {
	items []interface{}
}

type QueueCursor struct {
	Value interface{}
	q     *Queue
	i     int
}

func NewQueue() *Queue {
	return &Queue{
		items: []interface{}{},
	}
}

func (me *Queue) Enqueue(item interface{}) {
	me.items = append(me.items, item)
}

func (me *Queue) Dequeue() interface{} {
	if me.IsEmpty() {
		return nil
	}

	ret := me.items[0]
	me.items = me.items[1:]

	return ret
}

func (me *Queue) Size() int {
	return len(me.items)
}

func (me *Queue) IsEmpty() bool {
	return len(me.items) == 0
}

func (me *Queue) First() IterableCursor {
	if me.IsEmpty() {
		return nil
	}

	return &QueueCursor{
		Value: me.items[0],
		i:     0,
		q:     me,
	}
}

func (me *QueueCursor) Next() IterableCursor {
	i := me.i + 1

	if i >= len(me.q.items) {
		return nil
	}

	return &QueueCursor{
		Value: me.q.items[i],
		i:     i,
		q:     me.q,
	}
}
