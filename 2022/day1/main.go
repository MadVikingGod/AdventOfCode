package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	calories := parseCalories(input)

	max := 0
	for _, cal := range calories {
		if cal > max {
			max = cal
		}
	}
	println(max)

	top3 := make([]int, 3)
	for _, cal := range calories {
		if cal > top3[0] {
			top3 = append([]int{cal}, top3[:2]...)
		} else if cal > top3[1] {
			top3 = append(top3[:1], cal, top3[1])
		} else if cal > top3[2] {
			top3[2] = cal
		}
	}
	println(top3[0] + top3[1] + top3[2])
	// 208503 is too high
}

func parseCalories(input string) []int {
	var calories []int
	for _, elfString := range strings.Split(input, "\n\n") {
		calorie := 0
		for _, line := range strings.Split(elfString, "\n") {
			if line == "" {
				continue
			}
			c, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			calorie += c
		}
		calories = append(calories, calorie)
	}
	return calories
}
