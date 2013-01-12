package main

import (
	"fmt"
	"os"
	"strconv"
)

import (
	"github.com/meatballhat/goalgs4"
)

const USAGE string = `Usage: algs4-gcd <uint> <uint>`

func main() {
	if len(os.Args) < 3 {
		fmt.Println(USAGE)
		os.Exit(1)
	}

	p, err := strconv.ParseUint(os.Args[1], 10, 0)
	if err != nil {
		fmt.Println("Wherps: ", err)
		os.Exit(1)
	}

	q, err := strconv.ParseUint(os.Args[2], 10, 0)
	if err != nil {
		fmt.Println("Wherps: ", err)
		os.Exit(1)
	}

	fmt.Println(algs4.Gcd(p, q))
}
