package main

import (
	"fmt"
	"strings"
)

var input = "hepxcrrq"

func main() {
	fmt.Println(next(input))
	fmt.Println(next(next(input)))
}

// hasStraight returns true if the password contains a straight of three letters.
func hasStraight(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1]-1 && s[i] == s[i+2]-2 {
			return true
		}
	}
	return false
}

// noConfusing returns true if the password does not contain i, o, or l.
func noConfusing(s string) bool {
	return !strings.ContainsAny(s, "iol")
}

// twoDoubles returns true if the password contains two different non-overlapping pairs of letters.
func twoDoubles(s string) bool {
	var count int
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			count++
			i++
		}
	}
	return count >= 2
}

// increment increments the password by one.
func increment(s string) string {
	b := []byte(s)
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] == 'z' {
			b[i] = 'a'
		} else {
			b[i]++
			break
		}
	}
	return string(b)
}

// valid returns true if the password is valid.
func valid(s string) bool {
	return hasStraight(s) && noConfusing(s) && twoDoubles(s)
}

// next returns the next valid password.
func next(s string) string {
	for {
		s = increment(s)
		if valid(s) {
			return s
		}
	}
}
