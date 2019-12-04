package main

import "fmt"

func main() {
	min := 235741
	max := 706948

	test := [6]int{1, 1, 1, 1, 1, 1}
	fmt.Println(doesntDecrease(test), hasDouble(test))
	test = [6]int{0, 5, 4, 3, 2, 2}
	fmt.Println(doesntDecrease(test), hasDouble(test))
	test = [6]int{9, 8, 7, 3, 2, 1}
	fmt.Println(doesntDecrease(test), hasDouble(test))
	//702026
	test = [6]int{6, 2, 0, 2, 0, 7}
	fmt.Println(doesntDecrease(test), hasDouble(test))

	count := 0
	count2 := 0
	for i := min; i < max; i++ {
		digits := [6]int{}
		digits[0] = i % 10
		digits[1] = (i / (10)) % 10
		digits[2] = (i / (100)) % 10
		digits[3] = (i / (1000)) % 10
		digits[4] = (i / (10000)) % 10
		digits[5] = (i / (100000)) % 10
		if doesntDecrease(digits) && hasDouble(digits) {
			count++
		}
		if doesntDecrease(digits) && hasDoubleExtended(digits) {
			count2++
		}

	}
	fmt.Println(count)
	fmt.Println(count2)

}

func doesntDecrease(digits [6]int) bool {
	return digits[5] <= digits[4] &&
		digits[4] <= digits[3] &&
		digits[3] <= digits[2] &&
		digits[2] <= digits[1] &&
		digits[1] <= digits[0]
}
func hasDouble(digits [6]int) bool {
	return digits[5] == digits[4] ||
		digits[4] == digits[3] ||
		digits[3] == digits[2] ||
		digits[2] == digits[1] ||
		digits[1] == digits[0]
}

func hasDoubleExtended(digits [6]int) bool {
	first := digits[5] == digits[4] && digits[4] != digits[3]
	second := digits[5] != digits[4] && digits[4] == digits[3] && digits[3] != digits[2]
	third := digits[4] != digits[3] && digits[3] == digits[2] && digits[2] != digits[1]
	fourth := digits[3] != digits[2] && digits[2] == digits[1] && digits[1] != digits[0]
	fifth := digits[2] != digits[1] && digits[1] == digits[0]
	return first || second || third || fourth || fifth
}

//4997
