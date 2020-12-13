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
	fmt.Println(paths(input))
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
	lengths := map[int]int{
		len(l) - 1: 1,
	}

	for i := len(l) - 2; i >= 0; i-- {
		count := 0
		if i+1 < len(l) && l[i+1]-l[i] < 4 {
			count += lengths[i+1]
		}
		if i+2 < len(l) && l[i+2]-l[i] < 4 {
			count += lengths[i+2]
		}
		if i+3 < len(l) && l[i+3]-l[i] < 4 {
			count += lengths[i+3]
		}
		lengths[i] = count
	}

	return lengths[0]
}
