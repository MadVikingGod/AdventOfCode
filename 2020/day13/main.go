package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Did part1 using a calculator.
// the arrival time is start + (x - start/x) for the lowest x

// I did part2 using a calculator also.  Search for Chinese Remainder Theorem

func main() {
	// This is only part 2

	vals := []pair{}
	input := "19,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,x,x,x,643,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,17,13,x,x,x,x,23,x,x,x,x,x,x,x,509,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,29"
	// input = "7,13,x,x,59,x,31,19"
	for i, x := range strings.Split(input, ",") {
		val, err := strconv.Atoi(strings.Trim(x, ","))
		if err == nil {
			vals = append(vals, pair{val, val - (i % val)})
		}
	}
	fmt.Println(vals)
	start := (100000000000000/vals[0].val + 1) * vals[0].val
	// start = 7
	for i := start; true; i += vals[0].val {
		done := true
		for _, val := range vals {
			if !val.isValid(i) {
				done = false
				break
			}
		}
		if done {
			fmt.Println(i)
			return
		}
	}
}

type pair struct {
	val int
	pos int
}

func (p pair) isValid(i int) bool {
	return i%p.val == p.pos
}
