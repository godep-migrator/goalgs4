package main

import (
	"fmt"
	"os"
	"strconv"
)

import (
	"goalgs4"
)

const (
	USAGE = "Usage: algs4-rolls <nrolls>"
	SIDES = 6
)

func die(err error) {
	fmt.Println("Wat.", err)
	os.Exit(2)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(USAGE)
		os.Exit(1)
	}

	nrolls, err := strconv.Atoi(os.Args[1])
	if err != nil {
		die(err)
	}

	for _, roll := range algs4.Rolls(nrolls) {
		fmt.Println(roll)
	}
}
