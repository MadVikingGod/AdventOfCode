package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

//go:embed testInput.txt
var testInput string

func main() {
	signals := parseJsons(input)
	sum := 0
	for i := 0; i < len(signals)/2; i++ {
		cmp := compare(signals[i*2], signals[i*2+1])
		if cmp == 0 {
			panic("should not be equal")
		}
		if cmp == -1 {
			sum += i + 1
		}

	}

	// signals = parseJsons(in)

	fmt.Println(sum)

	two := []any{[]any{float64(2)}}
	six := []any{[]any{float64(6)}}
	signals = append(signals, two, six)
	sort.Slice(signals, func(i, j int) bool {
		return compare(signals[i], signals[j]) == -1
	})

	twoIndex := 0
	sixIndex := 0

	for i, signal := range signals {
		if compare(signal, two) == 0 {
			twoIndex = i + 1
		}
		if compare(signal, six) == 0 {
			sixIndex = i + 1
		}
	}
	fmt.Println(twoIndex * sixIndex)

}

func compare(left, right any) int {
	lInt, lok := left.(float64)
	rInt, rok := right.(float64)
	if lok && rok {
		return compareInts(int(lInt), int(rInt))
	}
	if !lok && !rok {
		//both are slices
		lSlice := left.([]any)
		rSlice := right.([]any)
		for i := 0; i < len(lSlice); i++ {
			if i >= len(rSlice) {
				return 1
			}
			cmp := compare(lSlice[i], rSlice[i])
			if cmp != 0 {
				return cmp
			}
		}
		if len(lSlice) < len(rSlice) {
			return -1
		}
		return 0
	}
	if lok {
		return compare([]any{left}, right)
	}
	return compare(left, []any{right})
}

func compareInts(left, right int) int {
	if left < right {
		return -1
	} else if left > right {
		return 1
	} else {
		return 0
	}
}

func parseJsons(input string) [][]any {
	output := [][]any{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		signal := []any{}
		err := json.Unmarshal([]byte(line), &signal)
		if err != nil {
			panic(err)
		}
		output = append(output, signal)
	}
	return output
}
