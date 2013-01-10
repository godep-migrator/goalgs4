package algs4

type Queue struct {
	items []interface{}
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
