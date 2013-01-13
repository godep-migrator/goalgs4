package algs4

type ArrayStack struct {
	items []interface{}
}

type ArrayStackCursor struct {
	v interface{}
	s *ArrayStack
	i int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		items: []interface{}{},
	}
}

func (me *ArrayStack) Push(item interface{}) {
	me.items = append(me.items, item)
}

func (me *ArrayStack) Pop() interface{} {
	if me.IsEmpty() {
		return nil
	}

	lastIdx := me.Size() - 1
	lastItem := me.items[lastIdx]
	me.items = me.items[:lastIdx]

	return lastItem
}

func (me *ArrayStack) IsEmpty() bool {
	return len(me.items) == 0
}

func (me *ArrayStack) Size() int {
	return len(me.items)
}

func (me *ArrayStack) First() IterableCursor {
	if me.IsEmpty() {
		return nil
	}

	last := len(me.items) - 1

	return &ArrayStackCursor{
		v: me.items[last],
		i: last,
		s: me,
	}
}

func (me *ArrayStackCursor) Value() interface{} {
	return me.v
}

func (me *ArrayStackCursor) Next() IterableCursor {
	i := me.i - 1

	if i < 0 {
		return nil
	}

	return &ArrayStackCursor{
		v: me.s.items[i],
		i: i,
		s: me.s,
	}
}
