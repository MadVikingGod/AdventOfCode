package main

import (
	"fmt"
	"sort"
)

func main() {
	input = append(input, 0)
	sort.Ints(input)
	counts := count(diff(input))
	fmt.Println(counts[1], counts[3]+1)
	fmt.Println(counts[1] * (counts[3] + 1))
}

func diff(l []int) []int {
	out := make([]int, len(l)-1)
	for i := 1; i < len(l); i++ {
		out[i-1] = l[i] - l[i-1]
	}
	return out
}

func count(l []int) map[int]int {
	out := map[int]int{}
	for _, i := range l {
		out[i] += 1
	}
	return out
}

func paths(l []int) int {
	count := 1
	length := len(l)
	l = append(l, 0, 0, 0, 0)
	for i := 0; i < length; i++ {
		x := l[i]
		for _, y := range l[i+1 : i+4] {
			if y-x <= 3 {
				count++
			}
		}
		count--
	}
	return count
}
