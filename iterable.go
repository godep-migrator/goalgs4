package algs4

type Iterable interface {
	First() IterableCursor
	IsEmpty() bool
	Size() int
}

type IterableCursor interface {
	Next() IterableCursor
	Value() interface{}
}
