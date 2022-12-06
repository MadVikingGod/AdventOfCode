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
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return true
			}
		}
	}
	return false
}
