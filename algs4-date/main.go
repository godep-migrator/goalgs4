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
	if len(os.Args) < 4 {
		die(errors.New("We need 3 ints!"))
	}

	m, err := strconv.Atoi(os.Args[1])
	if err != nil {
		die(err)
	}

	d, err := strconv.Atoi(os.Args[2])
	if err != nil {
		die(err)
	}

	y, err := strconv.Atoi(os.Args[3])
	if err != nil {
		die(err)
	}

	date := algs4.NewDate(m, d, y)
	fmt.Println(date)
}
