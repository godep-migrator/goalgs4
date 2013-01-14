package algs4

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
	if me.IsEmpty() {
		node := &Node{item: item}
		me.first = node
		me.last = node
		me.n++
		return
	}

	oldlast := me.last
	me.last = &Node{
		item: item,
	}
	oldlast.next = me.last
	me.n++
}

func (me *LinkedListQueue) Dequeue() interface{} {
	if me.IsEmpty() {
		return nil
	}

	oldfirst := me.first
	me.first = oldfirst.next
	oldfirst.next = nil

	me.n--

	if me.IsEmpty() {
		me.last = nil
	}

	return oldfirst.item
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

func (me *LinkedListQueueCursor) Value() interface{} {
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
