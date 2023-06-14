package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamples(t *testing.T) {
	input := `turn on 0,0 through 999,999
toggle 0,0 through 999,0
turn off 499,499 through 500,500`
	insts, _ := parseInstructions(input)
	g := grid{}

	g = insts[0](g)
	assert.Equal(t, 1000000, g.count())

	g = insts[1](g)
	assert.Equal(t, 999000, g.count())

	g = insts[2](g)
	assert.Equal(t, 998996, g.count())
}
