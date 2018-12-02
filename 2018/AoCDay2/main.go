package main

import (
	"fmt"
	"reflect"
)

func main() {

	twoCount := 0
	threeCount := 0

	for _, id := range Input {
		if hasTwo(id) {
			twoCount++
		}
		if hasThree(id) {
			threeCount++
		}
	}
	fmt.Println(threeCount * twoCount)
	fmt.Println(reflect.DeepEqual(Input, Input2))

	for i, id := range Input {
		for _, id2 := range Input[i+1:] {
			if distance(id, id2) == 1 {
				fmt.Printf("%s\n%s\n", id, id2)
			}
		}
	}
}

func distance(a, b string) int {
	d := 0
	B := []rune(b)
	for i, c := range a {
		if c != B[i] {
			d++
		}
	}
	return d
}

func hasTwo(in string) bool {
	found := map[rune]int{}
	for _, c := range in {
		found[c] = found[c] + 1
	}
	for _, v := range found {
		if v == 2 {
			return true
		}

	}
	return false
}
func hasThree(in string) bool {
	found := map[rune]int{}
	for _, c := range in {
		found[c] = found[c] + 1
	}
	for _, v := range found {
		if v == 3 {
			return true
		}

	}
	return false
}
