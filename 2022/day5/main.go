package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
)

/*
                    [Q]     [P] [P]
                [G] [V] [S] [Z] [F]
            [W] [V] [F] [Z] [W] [Q]
        [V] [T] [N] [J] [W] [B] [W]
    [Z] [L] [V] [B] [C] [R] [N] [M]
[C] [W] [R] [H] [H] [P] [T] [M] [B]
[Q] [Q] [M] [Z] [Z] [N] [G] [G] [J]
[B] [R] [B] [C] [D] [H] [D] [C] [N]
 1   2   3   4   5   6   7   8   9
*/

var spaces = [][]byte{
	{},
	[]byte("BQC"),
	[]byte("RQWZ"),
	[]byte("BMRLV"),
	[]byte("CZHVTW"),
	[]byte("DZHBNVG"),
	[]byte("HNPCJFVQ"),
	[]byte("DGTRWZS"),
	[]byte("CGMNBWZP"),
	[]byte("NJBMWQFP"),
}

//go:embed input.txt
var input string

func main() {
	var instructions []instruction
	for _, s := range strings.Split(input, "\n") {
		if s == "" {
			continue
		}
		instructions = append(instructions, praseInstruction(s))
	}
	// Copy for part2
	spaces2 := [][]byte{}
	for _, s := range spaces {
		spaces2 = append(spaces2, append([]byte{}, s...))
	}

	for _, i := range instructions {
		spaces = move(i, spaces)
	}
	output := bytes.Buffer{}
	for _, s := range spaces {
		if len(s) == 0 {
			continue
		}
		last := s[len(s)-1]
		output.WriteByte(last)
	}
	fmt.Println(output.String())

	for _, i := range instructions {
		spaces2 = move2(i, spaces2)
	}
	output = bytes.Buffer{}
	for _, s := range spaces2 {
		if len(s) == 0 {
			continue
		}
		last := s[len(s)-1]
		output.WriteByte(last)
	}
	fmt.Println(output.String())
}

type instruction struct {
	count    int
	from, to int
}

func praseInstruction(s string) instruction {
	var i instruction
	fmt.Sscanf(s, "move %d from %d to %d", &i.count, &i.from, &i.to)
	return i
}

func move(i instruction, spaces [][]byte) [][]byte {
	var (
		from = spaces[i.from]
		to   = spaces[i.to]
	)
	rev := reverse(from[len(from)-i.count:])
	to = append(to, rev...)
	from = from[:len(from)-i.count]
	spaces[i.from] = from
	spaces[i.to] = to
	return spaces
}

func reverse(b []byte) []byte {
	for i := len(b)/2 - 1; i >= 0; i-- {
		opp := len(b) - 1 - i
		b[i], b[opp] = b[opp], b[i]
	}
	return b
}

func move2(i instruction, spaces [][]byte) [][]byte {
	var (
		from = spaces[i.from]
		to   = spaces[i.to]
	)
	to = append(to, from[len(from)-i.count:]...)
	from = from[:len(from)-i.count]
	spaces[i.from] = from
	spaces[i.to] = to
	return spaces
}
