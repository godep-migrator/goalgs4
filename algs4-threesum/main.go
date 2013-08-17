package main

import (
	"fmt"
	"log"
	"os"

	. "goalgs4"
)

func main() {
	fd, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	a, err := ReadInts(fd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ThreeSumCount(a))
}
