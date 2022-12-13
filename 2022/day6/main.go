package main

import (
	_ "embed"
	"math/bits"
)

//go:embed input.txt
var input string

func main() {
	println(findStart(input))
	println(findStartOfMessage(input))
}

// findStart returns the position of the first non-repeating 4 character substring
func findStart(s string) int {
	for i := 3; i < len(s); i++ {
		if s[i-3] != s[i-2] && s[i-3] != s[i-1] && s[i-3] != s[i] &&
			s[i-2] != s[i-1] && s[i-2] != s[i] &&
			s[i-1] != s[i] {
			return i + 1
		}
	}
	return -1
}

// findStartofMessage returns the position of the first non-repeating 14 character substring
func findStartOfMessage(s string) int {
	var set uint = 0
	for i := 0; i < len(s); i++ {
		set ^= 1 << (s[i] - 'a')
		if i >= 13 {
			set ^= 1 << (s[i-13] - 'a')
		}
		if bits.OnesCount(set) == 13 {
			return i + 2
		}
	}
	return -1
}
