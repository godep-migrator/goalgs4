package main

import (
	"fmt"
	"os"
	"strconv"
)

import (
	"github.com/meatballhat/box-o-sand/algs4/src/go/algs4"
)

const (
	USAGE string = "Usage: algs4-rolls <nrolls>"
	SIDES        = 6
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(USAGE)
		os.Exit(1)
	}

	nrolls, err := strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		fmt.Println("Wat.", err)
		os.Exit(2)
	}

	rolls, err := algs4.Rolls(nrolls)

	for i := 1; i <= algs4.ROLL_SIDES; i += 1 {
		fmt.Println(rolls[i])
	}
}
