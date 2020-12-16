package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	input := map[int]struct {
		prev int
		last int
	}{
		16: {0, 1}, 1: {0, 2}, 0: {0, 3}, 18: {0, 4}, 12: {0, 5}, 14: {0, 6}, 19: {0, 7},
	}
	last := 19

	// input = map[int]struct {
	// 	prev int
	// 	last int
	// }{
	// 	0: {0, 1}, 3: {0, 2}, 6: {0, 3},
	// }
	// last = 6
	// input = map[int]struct {
	// 	prev int
	// 	last int
	// }{
	// 	1: {0, 1}, 3: {0, 2}, 2: {0, 3},
	// }
	// last = 2
	// input = map[int]struct {
	// 	prev int
	// 	last int
	// }{
	// 	2: {0, 1}, 1: {0, 2}, 3: {0, 3},
	// }
	// last = 6
	for i := len(input) + 1; i <= 30000000; i++ {

		next := 0
		if input[last].prev != 0 {
			next = input[last].last - input[last].prev
		}
		val := input[next]
		val.prev = val.last
		val.last = i
		input[next] = val
		last = next
		if i == 2020 {
			fmt.Println(i, next)
		}

	}
	fmt.Println(last)
	fmt.Println(time.Now().Sub(now))
}
