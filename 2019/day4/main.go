package main

import (
	"fmt"
	"strconv"
)

func main() {
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

}

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
	first := num[0] == num[1] && num[1] != num[2]
	second := num[0] != num[1] && num[1] == num[2] && num[2] != num[3]
	third := num[1] != num[2] && num[2] == num[3] && num[3] != num[4]
	fourth := num[2] != num[3] && num[3] == num[4] && num[4] != num[5]
	fifth := num[3] != num[4] && num[4] == num[5]
	return first || second || third || fourth || fifth
}

//4997
