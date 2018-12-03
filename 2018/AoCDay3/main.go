package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/MadVikingGod/AdventOfCode/2018/helpers"
)

// #1 @ 82,901: 26x12
// #(\d+) @ (\d+),(\d+): (\d+)x(\d+)

var regX = regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)

type claim struct {
	number int
	x      int
	y      int
	width  int
	height int
}

func main() {

	fabric := make([][]int, 1000)
	for i := range fabric {
		fabric[i] = make([]int, 1000)
	}

	input, err := helpers.GetInput(3)
	if err != nil {
		log.Panic(err)
	}

	//log.Print(input)
	claims := []claim{}

	for _, i := range input {
		claims = append(claims, parseClaim(i))
	}

	for _, c := range claims {
		c.Mark(fabric)
	}

	count := 0

	for _, x := range fabric {
		for _, c := range x {
			if c > 1 {
				count++
			}
		}
	}
	fmt.Println(count)

	for _, c := range claims {
		if !c.IsOverlap(fabric) {
			fmt.Println(c.number)
		}
	}
}

func (c claim) Mark(fabric [][]int) {
	for x := c.x; x < c.x+c.width; x++ {
		for y := c.y; y < c.y+c.height; y++ {
			fabric[x][y]++
		}
	}
}

func (c claim) IsOverlap(fabric [][]int) bool {
	for x := c.x; x < c.x+c.width; x++ {
		for y := c.y; y < c.y+c.height; y++ {
			if fabric[x][y] > 1 {
				return true
			}
		}
	}
	return false
}

func parseClaim(input string) claim {

	fields := strings.FieldsFunc(input, func(c rune) bool { return !unicode.IsNumber(c) })
	number, _ := strconv.Atoi(fields[0])
	x, _ := strconv.Atoi(fields[1])
	y, _ := strconv.Atoi(fields[2])
	width, _ := strconv.Atoi(fields[3])
	height, _ := strconv.Atoi(fields[4])
	return claim{number, x, y, width, height}
}
