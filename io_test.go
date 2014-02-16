package algs4_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/meatballhat/goalgs4"
)

const intsString = `
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
	ints, err := algs4.ReadInts(strings.NewReader(intsString))

	if err != nil {
		t.Errorf("Wrong wrong wrong: %v", err)
		return
	}

	if ints == nil {
		t.Errorf("Nothing in the ints!: %v", ints)
		return
	}

	fmt.Println("ints =", ints)

	if ints[0] != 999 {
		t.Errorf("fail on 0: 999 != %v", ints[0])
		return
	}

	if ints[4] != 71878 {
		t.Errorf("fail on 4: 71878 != %v", ints[4])
		return
	}

	if ints[6] != 0 {
		t.Errorf("fail on 6: 0 != %v", ints[6])
		return
	}
}
