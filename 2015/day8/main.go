package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	count := 0
	count2 := 0
	for _, line := range strings.Split(input, "\n") {
		count += len(line) - lenLine(line)
		count2 += encLen(line) - len(line)
	}
	println(count)
	println(count2)
}

func lenLine(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' {
			i++
			if s[i] == 'x' {
				i += 2
			}
		}
		count++
	}
	return count - 2
}

func encLen(s string) int {
	count := 2
	for _, c := range s {
		if c == '\\' || c == '"' {
			count++
		}
		count++
	}
	return count
}
