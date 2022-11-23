package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	boxes := boxFromInput(input)
	sum := 0
	ribbon := 0
	for _, box := range boxes {
		sum += box.wrappingPaper()
		ribbon += box.ribbon()
	}
	fmt.Println(sum)
	fmt.Println(ribbon)
}

// box is a struct that represents a box
type box struct {
	l, w, h int
}

// wrappingPaper calculates the amount of wrapping paper needed for a box
func (b box) wrappingPaper() int {
	return 2*b.l*b.w + 2*b.w*b.h + 2*b.h*b.l + min(b.l*b.w, b.w*b.h, b.h*b.l)
}

// ribbon calculates the amount of ribbon needed for a box
func (b box) ribbon() int {
	return 2*(b.l+b.w+b.h-max(b.l, b.w, b.h)) + b.l*b.w*b.h
}

func min(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func max(nums ...int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// boxFromInput takes a string and returns a box
func boxFromInput(input string) []box {
	boxes := []box{}
	for _, line := range strings.Split(input, "\n") {
		dimensions := strings.Split(line, "x")
		l, _ := strconv.Atoi(dimensions[0])
		w, _ := strconv.Atoi(dimensions[1])
		h, _ := strconv.Atoi(dimensions[2])
		boxes = append(boxes, box{l, w, h})
	}
	return boxes
}
