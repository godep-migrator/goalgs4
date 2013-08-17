package main

import (
	"fmt"
	"os"
	"sort"

	. "goalgs4"
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

	whiteList64, err := ReadInts(whiteListFile)
	if err != nil {
		fmt.Println("UGH: ", err)
		os.Exit(1)
	}

	whiteList := []int{}

	for _, i64 := range whiteList64 {
		whiteList = append(whiteList, int(i64))
	}

	sort.Ints(whiteList)

	inputs, err := ReadInts(os.Stdin)
	if err != nil {
		fmt.Println("Failed to read inputs:", err)
		os.Exit(1)
	}

	for _, key := range inputs {
		i := int(key)
		if BinarySearchRank(i, whiteList) == -1 {
			fmt.Println(i)
		}
	}
}
