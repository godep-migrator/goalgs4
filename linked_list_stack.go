package algs4

type LinkedListStack struct {
	first *Node
	n     int
}

type LinkedListStackCursor struct {
	n *Node
	s *LinkedListStack
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{}
}

func (me *LinkedListStack) Push(item interface{}) {
	oldfirst := me.first
	me.first = &Node{
		item: item,
		next: oldfirst,
	}
	me.n++
}

func (me *LinkedListStack) Pop() interface{} {
	if me.IsEmpty() {
		return nil
	}

	item := me.first.item
	me.first = me.first.next
	me.n--

	return item
}

func (me *LinkedListStack) IsEmpty() bool {
	return me.first == nil
}

func (me *LinkedListStack) Size() int {
	return me.n
}

func (me *LinkedListStack) First() IterableCursor {
	if me.IsEmpty() {
		return nil
	}

	return &LinkedListStackCursor{
		n: me.first,
		s: me,
	}
}

func (me *LinkedListStackCursor) Value() interface{} {
	return me.n.Value()
}

func (me *LinkedListStackCursor) Next() IterableCursor {
	next := me.n.next

	if next == nil {
		return nil
	}

	return &LinkedListStackCursor{
		n: next,
		s: me.s,
	}
}
