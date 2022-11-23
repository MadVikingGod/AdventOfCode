package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var things = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func main() {
	sues := parseSues(input)

	fmt.Println(findMatch(sues, things))
	fmt.Println(findMatchPart2(sues, things))
}

// Sue is a struct that represents a Sue
type Sue struct {
	number int
	things map[string]int
}

var reg = regexp.MustCompile(`Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)`)

func parseSues(input string) []Sue {
	sues := []Sue{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		match := reg.FindStringSubmatch(line)
		number, _ := strconv.Atoi(match[1])
		amount1, _ := strconv.Atoi(match[3])
		amount2, _ := strconv.Atoi(match[5])
		amount3, _ := strconv.Atoi(match[7])

		sues = append(sues, Sue{
			number: number,
			things: map[string]int{
				match[2]: amount1,
				match[4]: amount2,
				match[6]: amount3,
			},
		})
	}
	return sues
}

func findMatch(sues []Sue, things map[string]int) int {
	for _, sue := range sues {
		match := true
		for thing, amount := range sue.things {
			if things[thing] != amount {
				match = false
				break
			}
		}
		if match {
			return sue.number
		}
	}
	return -1
}

func findMatchPart2(sues []Sue, things map[string]int) int {
	for _, sue := range sues {
		match := true
		for thing, amount := range sue.things {
			if thing == "cats" || thing == "trees" {
				if things[thing] >= amount {
					match = false
					break
				}
			} else if thing == "pomeranians" || thing == "goldfish" {
				if things[thing] <= amount {
					match = false
					break
				}
			} else if things[thing] != amount {
				match = false
				break
			}
		}
		if match {
			return sue.number
		}
	}
	return -1
}
