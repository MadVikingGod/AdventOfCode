package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	r0 := parseRule(0)
	count := 0
	for _, msg := range input {
		for _, r := range r0 {
			if msg == r {
				count++
			}
		}
	}

	fmt.Println(count)

	rules[8] = "42 | 42 42 | 42 42 42"
	// rules[11] = "42 31 | 42 42 31 31 | 42 42 42 31 31 31"

	short := 0
	for _, s := range parseRule(0) {
		if short < len(s) {
			short = len(s)
		}
	}
	fmt.Println(short)
	short = 0
	for _, s := range input {
		if short < len(s) {
			short = len(s)
		}
	}
	fmt.Println(short)

}

func parseRule(i int) []string {
	if rules[i] == `"a"` {
		return []string{"a"}
	}
	if rules[i] == `"b"` {
		return []string{"b"}
	}
	out := []string{}
	str := []string{""}
	for _, token := range strings.Split(rules[i], " ") {
		if token == "|" {
			out = append(out, str...)
			str = []string{""}
			continue
		}
		num, _ := strconv.Atoi(token)

		tmp := []string{}
		for _, r := range parseRule(num) {
			for _, s := range str {
				tmp = append(tmp, s+r)
			}
		}
		str = tmp
	}

	return append(out, str...)
}

func parseRegex(i int) string {
	if rules[i] == `"a"` {
		return "a"
	}
	if rules[i] == `"b"` {
		return "b"
	}
	out := strings.Builder{}
	out.WriteString("(?:")
	for _, token := range strings.Split(rules[i], " ") {
		if token == "|" {
			out.WriteString("|")
		}
		num, _ := strconv.Atoi(token)
		out.WriteString(parseRegex(num))
	}
	out.WriteString(")")
	return  out.String()
}