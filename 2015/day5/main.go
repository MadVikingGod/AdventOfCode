package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	niceCount := 0
	extraNiceCount := 0
	for _, line := range strings.Split(input, "\n") {
		if isNice(line) {
			niceCount++
		}
		if hasDoubleDoubleLeter(line) && hasDoubleLetterWithOneBetween(line) {
			extraNiceCount++
		}
	}
	fmt.Println(niceCount)
	fmt.Println(extraNiceCount)
}

func isNice(s string) bool {
	return hasThreeVowels(s) && hasDoubleLetter(s) && !hasBadSubstring(s)
}

func hasThreeVowels(s string) bool {
	count := strings.Count(s, "a") + strings.Count(s, "e") + strings.Count(s, "i") + strings.Count(s, "o") + strings.Count(s, "u")
	return count >= 3
}

func hasDoubleLetter(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func hasBadSubstring(s string) bool {
	return strings.Contains(s, "ab") || strings.Contains(s, "cd") || strings.Contains(s, "pq") || strings.Contains(s, "xy")
}

func hasDoubleDoubleLeter(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if strings.Contains(s[i+2:], s[i:i+2]) {
			return true
		}
	}
	return false
}
func hasDoubleLetterWithOneBetween(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}
