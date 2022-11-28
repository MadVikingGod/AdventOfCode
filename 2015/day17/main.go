package main

import "fmt"

var input = []int{43, 3, 4, 10, 21, 44, 4, 6, 47, 41, 34, 17, 17, 44, 36, 31, 46, 9, 27, 38}

func main() {
	fmt.Println(countCaps(input, 150, 0))
	fmt.Println(lengths)
}

var lengths = map[int]int{}

func countCaps(input []int, cap int, l int) int {
	if cap == 0 {
		lengths[l]++
		return 1
	}
	if cap < 0 {
		return 0
	}
	count := 0
	for i, val := range input {
		count += countCaps(input[i+1:], cap-val, l+1)
	}
	return count
}
