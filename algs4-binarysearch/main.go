package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

import (
	"github.com/meatballhat/algs4"
)

const USAGE string = `Usage: algs4-binarysearch <whitelist-file>

You must also provide input on stdin.`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(USAGE)
		os.Exit(2)
	}

	whiteListFile, err := os.Open(string(os.Args[1]))
	if err != nil {
		fmt.Println("Failed to open whitelist file:", err)
	}

	whiteList, err := algs4.ReadInts(bufio.NewReader(whiteListFile))
	if err != nil {
		fmt.Println("UGH: ", err)
		os.Exit(1)
	}

	sort.Ints(whiteList)

	inputs, err := algs4.ReadInts(bufio.NewReader(os.Stdin))
	if err != nil {
		fmt.Println("Failed to read inputs:", err)
		os.Exit(1)
	}

	for _, key := range inputs {
		if algs4.BinarySearchRank(key, whiteList) == -1 {
			fmt.Println(key)
		}
	}
}
