package algs4

import (
	"reflect"
)

type Iterable interface {
	First() IterableCursor
	IsEmpty() bool
	Size() int
}

type IterableCursor interface {
	Next() IterableCursor
	Value() reflect.Value
}
