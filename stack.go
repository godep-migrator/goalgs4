package algs4

type Stack interface {
	Iterable
	Push(item interface{})
	Pop() interface{}
}
