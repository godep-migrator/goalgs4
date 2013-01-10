package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/meatballhat/box-o-sand/algs4/src/go/algs4"
)

func die(err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		die(errors.New("We need an int!"))
	}

	nValues, err := strconv.Atoi(os.Args[1])
	if err != nil {
		die(err)
	}

	a := algs4.NewAccumulator()
	for t := 0; t < nValues; t++ {
		a.AddDataValue(algs4.Random())
	}

	fmt.Println(a)
	return
}
