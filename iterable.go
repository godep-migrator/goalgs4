package algs4

type Iterable interface {
	First() IterableCursor
}

type IterableCursor interface {
	Next() IterableCursor
}
