package main

import "fmt"

func main() {
	fmt.Println(countIncreases(inputs))
	fmt.Println(countWindowIncreases(inputs, 3))
}

func countIncreases(list []int) int {
	count := 0

	for i := 1; i < len(list); i++ {
		if list[i-1] < list[i] {
			count++
		}
	}
	return count
}

func countWindowIncreases(list []int, windowSize int) int {
	count := 0
	for i := windowSize + 1; i <= len(list); i++ {
		w1 := sum(list[i-windowSize-1 : i-1])
		w2 := sum(list[i-windowSize : i])
		if w1 < w2 {
			count++
		}
	}
	return count
}

func sum(list []int) int {
	a := 0
	for i := 0; i < len(list); i++ {
		a += list[i]
	}
	return a
}
