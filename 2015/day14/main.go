package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	reindeers := parseReindeers(input)
	fmt.Println(reindeers)
	d := 0
	for _, r := range reindeers {
		d = max(d, r.distance(2503))
	}
	fmt.Println(d)

	// 1060 is too high
	fmt.Println(score(reindeers, 2503))
}

type reindeer struct {
	name    string
	speed   int
	stamina int
	rest    int
}

func (r reindeer) distance(t int) int {
	// 1. how many cycles of stamina + rest?
	cycles := t / (r.stamina + r.rest)
	// 2. how many seconds left over?
	remainder := t % (r.stamina + r.rest)
	// 3. how many seconds of stamina?
	stamina := min(r.stamina, remainder)
	// 4. how many seconds of rest?
	// _ := max(0, remainder-r.stamina)
	// 5. how far did we go?
	return r.speed * (cycles*r.stamina + stamina)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseReindeers(input string) []reindeer {
	reinders := []reindeer{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		var name string
		var speed, stamina, rest int
		_, err := fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &speed, &stamina, &rest)
		if err != nil {
			panic(err)
		}
		reinders = append(reinders, reindeer{name, speed, stamina, rest})
	}
	return reinders
}

func score(reindeers []reindeer, maxDistance int) int {
	score := make([]int, len(reindeers))
	winners := make([]int, 0, len(reindeers))
	for i := 1; i < maxDistance; i++ {
		max := 0
		for j, r := range reindeers {
			d := r.distance(i)
			if d > max {
				max = d
				winners = winners[:0]
			}
			if d == max {
				winners = append(winners, j)
			}
		}
		for _, w := range winners {
			score[w]++
		}
	}

	max := 0
	for _, s := range score {
		if s > max {
			max = s
		}
	}
	return max
}
