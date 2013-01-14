package algs4

import (
	"reflect"
)

type ArrayBag struct {
	items []interface{}
}

type ArrayBagCursor struct {
	v interface{}
	b *ArrayBag
	i int
}

func NewArrayBag() Bag {
	return &ArrayBag{
		items: []interface{}{},
	}
}

func (me *ArrayBag) Add(item interface{}) {
	me.items = append(me.items, item)
}

func (me *ArrayBag) IsEmpty() bool {
	return len(me.items) == 0
}

func (me *ArrayBag) Size() int {
	return len(me.items)
}

func (me *ArrayBag) First() IterableCursor {
	if me.IsEmpty() {
		return nil
	}

	return &ArrayBagCursor{
		v: me.items[0],
		i: 0,
		b: me,
	}
}

func (me *ArrayBagCursor) Value() reflect.Value {
	return reflect.ValueOf(me.v)
}

func (me *ArrayBagCursor) Next() IterableCursor {
	i := me.i + 1

	if i > me.b.Size()-1 {
		return nil
	}

	return &ArrayBagCursor{
		v: me.b.items[i],
		i: i,
		b: me.b,
	}
}
