package main

import (
    "fmt"
    "log"

    . "github.com/meatballhat/goalgs4"
)

func main() {
    s := NewLinkedListStack()

    for !Stdin.IsEmpty() {
        item, err := Stdin.ReadString()
        if err != nil {
            log.Fatal(err)
        }

        if item != "-" {
            s.Push(item)
        } else if len(item) > 0 {
            fmt.Printf("%v ", s.Pop())
        }
    }

    fmt.Printf("(%v left on stack)\n", s.Size())
}
