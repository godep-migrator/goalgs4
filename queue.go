package algs4

type Queue interface {
	Iterable
	Enqueue(item interface{})
	Dequeue() interface{}
}
