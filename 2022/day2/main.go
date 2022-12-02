package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	maxCount := 0
	for _, selfMap := range selfMaps {
		count := match(input, selfMap)
		if count > maxCount {
			maxCount = count
		}
	}
	println(maxCount)

	println(match2(input))
}

func match(input string, selfMap map[string]string) int {

	count := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		opp := opponentMap[parts[0]]
		self := selfMap[parts[1]]
		count += round(opp, self)
	}
	return count
}

func match2(input string) int {
	transform := map[string]map[string]int{
		"X": { //Lose
			"A": 3,
			"B": 1,
			"C": 2,
		},
		"Y": { //Draw
			"A": 1 + 3,
			"B": 2 + 3,
			"C": 3 + 3,
		},
		"Z": { //Win
			"A": 2 + 6,
			"B": 3 + 6,
			"C": 1 + 6,
		},
	}
	count := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		count += transform[parts[1]][parts[0]]
	}
	return count
}

func round(opp, self string) int {
	// Win is worth 6 points
	// Draw is worth 3 points
	// Loss is worth 0 points
	// R is worth 1 point
	// P is worth 2 points
	// S is worth 3 points
	throw := map[string]int{
		"R": 1,
		"P": 2,
		"S": 3,
	}
	if opp == self {
		return 3 + throw[self]
	}
	if (opp == "R" && self == "P") || (opp == "P" && self == "S") || (opp == "S" && self == "R") {
		return 6 + throw[self]
	}
	return 0 + throw[self]

}

// tables
var opponentMap = map[string]string{
	"A": "R",
	"B": "P",
	"C": "S",
}

var selfMaps = []map[string]string{
	{
		"X": "R",
		"Y": "P",
		"Z": "S",
	},
	{
		"X": "R",
		"Y": "S",
		"Z": "P",
	},
	{
		"X": "P",
		"Y": "R",
		"Z": "S",
	},
	{
		"X": "P",
		"Y": "S",
		"Z": "R",
	},
	{
		"X": "S",
		"Y": "R",
		"Z": "P",
	},
	{
		"X": "S",
		"Y": "P",
		"Z": "R",
	},
}
