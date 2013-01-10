package algs4

import (
	"fmt"
	"reflect"
)

type Bag struct {
	items map[string]interface{}
}

type Item struct {
	Name string
}

func NewBag() *Bag {
	return &Bag{
		items: map[string]interface{}{},
	}
}

func (me *Bag) Add(item interface{}) {
	me.items[me.hashItem(item)] = item
}

func (me *Bag) IsEmpty() bool {
	return len(me.items) == 0
}

func (me *Bag) Size() int {
	return len(me.items)
}

func (me *Bag) Each() <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)
		for _, item := range me.items {
			out <- item
		}
	}()

	return (<-chan interface{})(out)
}

func (me *Bag) hashItem(item interface{}) string {
	t := reflect.TypeOf(item)
	return fmt.Sprintf("%s%s%v", t.Name(), t.String(), &t)
}
