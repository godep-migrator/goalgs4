package algs4_test

import (
	"fmt"
	"strings"
	"testing"
)

import (
	"github.com/meatballhat/algs4"
)

const INTS_STRING = `
999
848
-7271
7384
71878
92
0
0
0
-99
`

func TestReadInts(t *testing.T) {
	ints, err := algs4.ReadInts(strings.NewReader(INTS_STRING))

	if err != nil {
		t.Error("Wrong wrong wrong: ", err)
	}

	if ints == nil {
		t.Error("Nothing in the ints!: ", ints)
	}

	fmt.Println("ints =", ints)

	if ints[0] != 999 {
		t.Error("fail on 0: 999 !=", ints[0])
	}

	if ints[4] != 71878 {
		t.Error("fail on 4: 71878 !=", ints[4])
	}

	if ints[6] != 0 {
		t.Error("fail on 6: 0 !=", ints[6])
	}
}
