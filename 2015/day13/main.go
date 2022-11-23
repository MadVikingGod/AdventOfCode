package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	h := parse(input)
	fmt.Println(h.findBestSeating())
	fmt.Println(h.findBestSeatingWithSelf())
}

// happness is a weighted graph of strings
type happness map[string]map[string]int

// parse parses the input and returns a happness graph
func parse(input string) happness {
	h := make(happness)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		var a, b, verb string
		var v int
		fmt.Sscanf(line, "%s would %s %d happiness units by sitting next to %s.", &a, &verb, &v, &b)
		b = strings.Trim(b, ".")
		if verb == "lose" {
			v = -v
		}
		if _, ok := h[a]; !ok {
			h[a] = make(map[string]int)
		}
		h[a][b] = v
	}
	return h
}

// getScore returns the score of the seating arrangement
func (h happness) getScore(seating []string) int {
	score := 0
	for i, a := range seating {
		b := seating[(i+1)%len(seating)]
		x := h[a][b]
		y := h[b][a]
		score += x
		score += y
	}
	return score
}

// permute returns all permutations of the input slice
func permute(input []string) [][]string {
	if len(input) == 1 {
		return [][]string{input}
	}
	var output [][]string
	for i, v := range input {
		rest := append([]string{}, input[:i]...)
		rest = append(rest, input[i+1:]...)
		for _, p := range permute(rest) {
			output = append(output, append([]string{v}, p...))
		}
	}
	return output
}

// getPeople returns a slice of all people in the graph
func (h happness) getPeople() []string {
	var people []string
	for p := range h {
		people = append(people, p)
	}
	return people
}

// findBestSeating returns the best seating arrangement
func (h happness) findBestSeating() ([]string, int) {
	var bestSeating []string
	var bestScore int
	for _, seating := range permute(h.getPeople()) {
		score := h.getScore(seating)
		if score > bestScore {
			bestScore = score
			bestSeating = seating
		}
	}
	return bestSeating, bestScore
}

// findBestSeatingWithSelf returns the best seating arrangement including self
func (h happness) findBestSeatingWithSelf() ([]string, int) {
	var bestSeating []string
	var bestScore int
	people := append(h.getPeople(), "self")
	for _, seating := range permute(people) {
		score := h.getScore(seating)
		if score > bestScore {
			bestScore = score
			bestSeating = seating
		}
	}
	return bestSeating, bestScore
}
