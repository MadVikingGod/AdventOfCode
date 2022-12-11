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

// total mod for this input is 9699690

//go:embed testInput.txt
var testInput string

// total mod for this input is 96577

func main() {
	pack := parsePack(input)

	for round := 0; round < 20; round++ {
		for _, monkey := range pack {
			monkey.turn(pack, 3, 9699690)
		}
	}

	fmt.Println(max2(pack))

	pack = parsePack(input)
	for round := 0; round < 10000; round++ {
		for _, monkey := range pack {
			monkey.turn(pack, 1, 9699690)
		}
	}

	fmt.Println(max2(pack))

}

func max2(pack []*monkey) int {
	var max1, max2 int
	for _, monkey := range pack {
		if monkey == nil {
			continue
		}
		if monkey.count > max1 {
			max1, max2 = monkey.count, max1
			continue
		}
		if monkey.count > max2 {
			max2 = monkey.count
		}
	}
	return max1 * max2
}

type monkey struct {
	items       []int
	operation   func(int) int
	divisor     int
	trueMonkey  int
	falseMonkey int

	count int
}

func (m *monkey) turn(pack []*monkey, div, mod int) {
	if m == nil {
		return
	}
	for _, item := range m.items {
		m.count++
		item = m.operation(item) / div % mod

		if item%m.divisor == 0 {
			pack[m.trueMonkey].items = append(pack[m.trueMonkey].items, item)
		} else {
			pack[m.falseMonkey].items = append(pack[m.falseMonkey].items, item)
		}
	}
	m.items = []int{}
}

var itemsRegex = regexp.MustCompile(`(\d+)`)

func parsePack(input string) []*monkey {
	out := [8]*monkey{}

	for _, m := range strings.Split(input, "\n\n") {
		monkey := &monkey{}
		var i int
		lines := strings.Split(m, "\n")
		fmt.Sscanf(lines[0], "Monkey %d:", &i)
		monkey.items = parseItems(lines[1])
		monkey.operation = parseOperation(lines[2])
		fmt.Sscanf(lines[3], "  Test: divisible by %d", &monkey.divisor)
		fmt.Sscanf(lines[4], "    If true: throw to monkey %d", &monkey.trueMonkey)
		fmt.Sscanf(lines[5], "    If false: throw to monkey %d", &monkey.falseMonkey)
		out[i] = monkey
	}

	return out[:]
}

func parseItems(input string) []int {
	out := []int{}
	for _, item := range itemsRegex.FindAllStringSubmatch(input, -1) {
		i, err := strconv.Atoi(item[1])
		if err != nil {
			panic(err)
		}
		out = append(out, i)
	}
	return out
}

func parseOperation(input string) func(int) int {
	input = input[23:]
	if input == "* old" {
		return func(i int) int { return i * i }
	}
	number, err := strconv.Atoi(input[2:])
	if err != nil {
		panic(err)
	}
	if input[0] == '+' {
		return func(i int) int { return i + number }
	}
	if input[0] == '*' {
		return func(i int) int { return i * number }
	}
	panic("unknown operation")
}
