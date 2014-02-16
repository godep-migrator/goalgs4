package main

import (
	"fmt"
	"log"

	. "github.com/meatballhat/goalgs4"
)

func main() {
	q := NewLinkedListQueue()

	for !Stdin.IsEmpty() {
		item, err := Stdin.ReadString()
		if err != nil {
			log.Fatal(err)
		}

		if item != "-" {
			q.Enqueue(item)
		} else if !q.IsEmpty() {
			fmt.Printf("%v ", q.Dequeue())
		}
	}

	fmt.Printf("(%v left on queue)\n", q.Size())
}
