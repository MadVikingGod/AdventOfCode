package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	rucksacks := PraseRucksacks(input)
	count := 0
	for _, r := range rucksacks {
		count += priority(r.GetRepeat())
	}
	println(count)

	count = 0
	for i := 0; i < len(rucksacks); i += 3 {
		count += priority(findBadge(rucksacks[i : i+3]))
	}
	println(count)
}

type rucksack string

func (r rucksack) GetRepeat() rune {
	part1 := r[0 : len(r)/2]
	part2 := r[len(r)/2:]

	for _, r := range part1 {
		if strings.ContainsRune(string(part2), r) {
			return r
		}
	}
	return 0
}

func PraseRucksacks(input string) []rucksack {
	rucksacks := []rucksack{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		rucksacks = append(rucksacks, rucksack(line))
	}
	return rucksacks
}

func priority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r) - int('a') + 1
	}
	return int(r) - int('A') + 27
}

func findBadge(rucksacks []rucksack) rune {
	for _, r := range rucksacks[0] {
		if strings.ContainsRune(string(rucksacks[1]), r) && strings.ContainsRune(string(rucksacks[2]), r) {
			return r
		}
	}
	return 0
}
