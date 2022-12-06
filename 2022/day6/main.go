package main

import _ "embed"

//go:embed input.txt
var input string

func main() {
	println(findStart(input))
	println(findStartOfMessage(input))
}

// findStart returns the position of the first non-repeating 4 character substring
func findStart(s string) int {
	for i := 3; i < len(s); i++ {
		if !hasRepeat(s[i-3 : i+1]) {
			return i + 1
		}
	}
	return -1
}

// findStartofMessage returns the position of the first non-repeating 14 character substring
func findStartOfMessage(s string) int {
	for i := 13; i < len(s); i++ {
		if !hasRepeat(s[i-13 : i+1]) {
			return i + 1
		}
	}
	return -1
}

func hasRepeat(s string) bool {
	set := make(map[rune]bool)
	for _, r := range s {
		if set[r] {
			return true
		}
		set[r] = true
	}
	return false
}
