package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	pairs := parsePairs(input)
	count := 0
	oCount := 0
	for _, pair := range pairs {
		if pair.first.contains(pair.second) || pair.second.contains(pair.first) {
			count++
		}
		if pair.first.overlap(pair.second) {
			oCount++
		}
	}
	println(count)
	println(oCount)

}

type section struct {
	min, max int
}

type pair struct {
	first, second section
}

func parsePairs(input string) []pair {
	pairs := []pair{}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ",")
		first := parseSection(parts[0])
		second := parseSection(parts[1])
		pairs = append(pairs, pair{first, second})
	}
	return pairs
}

func parseSection(input string) section {
	sides := strings.Split(input, "-")
	min, _ := strconv.Atoi(sides[0])
	max, _ := strconv.Atoi(sides[1])
	return section{min, max}
}

func (s section) contains(o section) bool {
	return s.min <= o.min && s.max >= o.max
}
func (s section) overlap(o section) bool {
	return s.min <= o.max && s.max >= o.min
}
