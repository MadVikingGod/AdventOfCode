package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	min := 235741
	max := 706948

	count := 0
	count2 := 0
	for i := min; i < max; i++ {
		num := strconv.Itoa(i)

		if doesntDecrease(num) && hasDouble(num) {
			count++
		}
		if doesntDecrease(num) && hasDoubleExtended(num) {
			count2++
		}

	}
	fmt.Println(count)
	fmt.Println(count2)
	fmt.Println(time.Now().Sub(start))
}

// Note: all of these checks would need a refactor if input was arbitrary length

func doesntDecrease(num string) bool {
	return num[0] <= num[1] &&
		num[1] <= num[2] &&
		num[2] <= num[3] &&
		num[3] <= num[4] &&
		num[4] <= num[5]
}
func hasDouble(num string) bool {
	return num[0] == num[1] ||
		num[1] == num[2] ||
		num[2] == num[3] ||
		num[3] == num[4] ||
		num[4] == num[5]
}

func hasDoubleExtended(num string) bool {
	// This was done because a number like 333445 is valid
	// What we are checking is for groups like ^112, 2112, or 211$ exist.
	// This is fine because the middle group (2112) can only be in 3 places
	// because the number is 6 digits long.
	first := num[0] == num[1] && num[1] != num[2]
	second := num[0] != num[1] && num[1] == num[2] && num[2] != num[3]
	third := num[1] != num[2] && num[2] == num[3] && num[3] != num[4]
	fourth := num[2] != num[3] && num[3] == num[4] && num[4] != num[5]
	fifth := num[3] != num[4] && num[4] == num[5]
	return first || second || third || fourth || fifth
}

//4997
