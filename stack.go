package algs4

type Stack struct {
	items []interface{}
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
