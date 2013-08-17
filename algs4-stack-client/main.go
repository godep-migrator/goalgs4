package main

import (
	"fmt"
	"log"
	"os"

	. "goalgs4"
)

func main() {
	var stack Stack
	if len(os.Args) > 1 && os.Args[1] == "ll" {
		stack = NewLinkedListStack()
	} else {
		stack = NewArrayStack()
	}

	for !Stdin.IsEmpty() {
		i64, err := Stdin.ReadInt()
		if err != nil {
			log.Fatal(err)
		}

		stack.Push(i64)
	}

	for cursor := stack.First(); cursor != nil; cursor = cursor.Next() {
		fmt.Println(cursor.Value().Int())
	}
}
