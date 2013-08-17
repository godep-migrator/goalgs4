package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	. "goalgs4"
)

func main() {
	n, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	max := float64(1000000.0)
	a := []int64{}

	for i := int64(0); i < n; i++ {
		a = append(a, int64(RandomUniform(-max, max)))
	}

	timer := NewStopwatch()
	cnt := ThreeSumCount(a)
	elapsed := timer.ElapsedTime()

	fmt.Printf("%v triples %v\n", cnt, elapsed)
}
