package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = "1113222113"

func main() {
	for i := 0; i < 40; i++ {
		input = LookAndSay(input)
	}
	fmt.Println(len(input))

	for i := 0; i < 10; i++ {
		input = LookAndSay(input)
	}
	fmt.Println(len(input))
}

// LookAndSay returns a string that counts the number of times a digit is repeated.
func LookAndSay(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		count := 1
		for i+1 < len(s) && s[i] == s[i+1] {
			count++
			i++
		}
		b.WriteString(strconv.Itoa(count))
		b.WriteByte(s[i])
	}
	return b.String()
}
