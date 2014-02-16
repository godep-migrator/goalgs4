package algs4_test

import (
	"testing"

	. "github.com/meatballhat/goalgs4"
)

func TestQueueIsIterable(t *testing.T) {
	q := NewArrayQueue()

	for i := 0; i < 1000; i += 10 {
		q.Enqueue(i)
	}

	iterations := 0
	for cursor := q.First(); cursor != nil; cursor = cursor.Next() {
		iterations++
	}

	if iterations != 100 {
		t.Errorf("iterations != 100: %v", iterations)
	}
}
