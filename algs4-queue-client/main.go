package main

import (
	"fmt"
	"log"
	"os"

	. "goalgs4"
)

var (
	usage = "Usage: %v <filename>\n"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
		os.Exit(1)
	}

	in, err := NewIn(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	q := NewArrayQueue()
	for !in.IsEmpty() {
		i64, err := in.ReadInt()
		if err != nil {
			break
		}

		q.Enqueue(i64)
	}

	n := q.Size()
	var a []int64

	for i := 0; i < n; i++ {
		v := q.Dequeue()
		if i64, ok := v.(int64); ok {
			a = append(a, i64)
		}
	}

	for _, i := range a {
		fmt.Println(i)
	}
}
