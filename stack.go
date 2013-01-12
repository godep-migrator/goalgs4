package algs4

type Stack struct {
	items []interface{}
}

type StackCursor struct {
	v interface{}
	s *Stack
	i int
}

func NewStack() *Stack {
	return &Stack{
		items: []interface{}{},
	}
}

func (me *Stack) Push(item interface{}) {
	me.items = append(me.items, item)
}

func (me *Stack) Pop() interface{} {
	if me.IsEmpty() {
		return nil
	}

	lastIdx := me.Size() - 1
	lastItem := me.items[lastIdx]
	me.items = me.items[:lastIdx]

	return lastItem
}

func (me *Stack) IsEmpty() bool {
	return len(me.items) == 0
}

func (me *Stack) Size() int {
	return len(me.items)
}

func (me *Stack) First() IterableCursor {
	if me.IsEmpty() {
		return nil
	}

	last := len(me.items) - 1

	return &StackCursor{
		v: me.items[last],
		i: last,
		s: me,
	}
}

func (me *StackCursor) Value() interface{} {
	return me.v
}

func (me *StackCursor) Next() IterableCursor {
	i := me.i - 1

	if i < 0 {
		return nil
	}

	return &StackCursor{
		v: me.s.items[i],
		i: i,
		s: me.s,
	}
}
