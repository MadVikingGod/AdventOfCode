package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func main() {
	start := time.Now()
	fmt.Println(part2(input))
	dur := time.Since(start)
	fmt.Printf("Time: %v\n", dur)
}

func part2(input string) int {
	parts := strings.Split(input, "\n\n")
	seedMap := parseSeedMap(parts[0])

	maps := []Map{}
	for _, m := range parts[1:] {
		maps = append(maps, parse_map(strings.Split(m, "\n")))
	}

	slices.Reverse(maps)

	for i := 0; ; i++ {
		seed := i
		for _, m := range maps {
			seed = m.nextLoc(seed)
		}
		if seedMap.isSeed(seed) {
			return i
		}
	}
}

type SeedMap struct {
	seeds []rng
}

type rng struct {
	start, len int
}

func (s SeedMap) isSeed(input int) bool {
	for _, seed := range s.seeds {
		if input >= seed.start && input < seed.start+seed.len {
			return true
		}
	}
	return false
}

func parseSeedMap(input string) SeedMap {
	parts := strings.Split(input, " ")
	seeds := []rng{}
	for i := 1; i < len(parts); i += 2 {
		start, _ := strconv.Atoi(parts[i])
		length, _ := strconv.Atoi(parts[i+1])
		seeds = append(seeds, rng{start, length})
	}
	return SeedMap{seeds}
}

type Map struct {
	parts []row
}
type row struct {
	inStart, outStart, len int
}

func parse_map(input []string) Map {
	m := Map{parts: []row{}}
	for _, line := range input[1:] {
		parts := strings.Split(line, " ")
		inStart, _ := strconv.Atoi(parts[0])
		outStart, _ := strconv.Atoi(parts[1])
		length, _ := strconv.Atoi(parts[2])
		m.parts = append(m.parts, row{inStart, outStart, length})
	}
	return m
}

func (m Map) nextLoc(input int) int {
	for _, part := range m.parts {
		if input >= part.inStart && input < part.inStart+part.len {
			return part.outStart + (input - part.inStart)
		}
	}
	return input
}

func (m Map) prevLoc(output int) int {
	for _, part := range m.parts {
		if output >= part.outStart && output < part.outStart+part.len {
			return part.inStart + (output - part.outStart)
		}
	}
	return output
}
