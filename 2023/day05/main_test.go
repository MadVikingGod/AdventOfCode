package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed test_input.txt
var testInput string

func TestParseMap(t *testing.T) {
	input := strings.Split(testInput, "\n\n")
	m := parse_map(strings.Split(input[1], "\n"))
	assert.Len(t, m.parts, 2)
	assert.Equal(t, Map{parts: []row{{50, 98, 2}, {52, 50, 48}}}, m)
	m = parse_map(strings.Split(input[2], "\n"))
	assert.Len(t, m.parts, 3)
	assert.Equal(t, Map{parts: []row{{0, 15, 37}, {37, 52, 2}, {39, 0, 15}}}, m)
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 46, part2(testInput))
}
