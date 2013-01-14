package algs4

type Queue interface {
	Enqueue(item interface{})
	Dequeue() interface{}
	Size() int
	IsEmpty() bool
	First() IterableCursor
}
