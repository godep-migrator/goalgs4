package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	. "github.com/meatballhat/goalgs4"
)

func main() {
	ops := NewArrayStack()
	vals := NewArrayStack()

	for !Stdin.IsEmpty() {
		s, err := Stdin.ReadString()
		if err != nil {
			log.Fatal(err)
		}

		switch s {
		case "(":
		case "+", "-", "*", "/", "sqrt":
			ops.Push(s)
		case ")":
			o := ops.Pop()
			v := vals.Pop()

			var (
				ok    bool
				op    string
				value float64
			)

			if op, ok = o.(string); !ok {
				log.Fatalf("Popped operator is not a string!: %v", op)
			}

			if value, ok = v.(float64); !ok {
				log.Fatalf("Popped value is not a float64!: %v", value)
			}

			switch op {
			case "+", "-", "*", "/":
				var nextValue float64

				nextv := vals.Pop()
				if nextValue, ok = nextv.(float64); !ok {
					log.Fatalf("Popped value is not a float64!: %v", nextValue)
				}

				switch op {
				case "+":
					value = nextValue + value
				case "-":
					value = nextValue - value
				case "*":
					value = nextValue * value
				case "/":
					value = nextValue / value
				}
			case "sqrt":
				value = math.Sqrt(value)
			}

			vals.Push(value)
		default:
			f64, err := strconv.ParseFloat(s, 64)
			if err != nil {
				log.Fatal(err)
			}

			vals.Push(f64)
		}
	}

	out := fmt.Sprintf("%.15f", vals.Pop())
	out = strings.TrimRight(out, "0")
	if strings.HasSuffix(out, ".") {
		fmt.Printf("%v0\n", out)
		return
	}

	fmt.Printf("%v\n", out)
}
