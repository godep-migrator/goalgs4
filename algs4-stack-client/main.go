package main

import (
	"fmt"
	"log"

	. "github.com/meatballhat/goalgs4"
)

func main() {
	stack := NewArrayStack()

	for !Stdin.IsEmpty() {
		i64, err := Stdin.ReadInt()
		if err != nil {
			log.Fatal(err)
		}

		stack.Push(i64)
	}

	for cursor := stack.First(); cursor != nil; cursor = cursor.Next() {
		if i64, ok := cursor.Value().(int64); ok {
			fmt.Println(i64)
		}
	}
}
