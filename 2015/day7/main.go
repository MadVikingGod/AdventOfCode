package main

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	circuit := parseCircuit(input)

	a := circuit["a"].eval(circuit)
	println(a)

	circuit = parseCircuit(input)
	circuit["b"] = immediate(a)
	println(circuit["a"].eval(circuit))

}

type net struct {
	val    int
	evaled bool
	f      func(map[string]*net) int
}

func (n *net) eval(cir map[string]*net) int {
	if !n.evaled {
		n.val = n.f(cir)
		n.evaled = true
	}
	return n.val
}

func val(v string) *net {
	return &net{f: func(cir map[string]*net) int {
		return cir[v].eval(cir)
	}}
}
func immediate(v int) *net {
	return &net{val: v, evaled: true}
}
func and(a, b string) *net {
	return &net{f: func(cir map[string]*net) int {
		return cir[a].eval(cir) & cir[b].eval(cir)
	}}
}
func or(a, b string) *net {
	return &net{f: func(cir map[string]*net) int {
		return cir[a].eval(cir) | cir[b].eval(cir)
	}}
}
func lshift(a string, b string) *net {
	return &net{f: func(cir map[string]*net) int {
		return cir[a].eval(cir) << cir[b].eval(cir)
	}}
}
func rshift(a string, b string) *net {
	return &net{f: func(cir map[string]*net) int {
		return cir[a].eval(cir) >> cir[b].eval(cir)
	}}
}
func not(a string) *net {
	return &net{f: func(cir map[string]*net) int {
		return ^cir[a].eval(cir)
	}}
}

var reTwoArg = regexp.MustCompile(`^(\w+) (\w+) (\w+) -> (\w+)$`)
var reNot = regexp.MustCompile(`^NOT (\w+) -> (\w+)$`)
var reVal = regexp.MustCompile(`^(\w+) -> (\w+)$`)

func parseCircuit(s string) map[string]*net {
	circuit := make(map[string]*net)
	for _, line := range strings.Split(s, "\n") {
		if matches := reTwoArg.FindStringSubmatch(line); matches != nil {
			if v, err := strconv.Atoi(matches[1]); err == nil {
				circuit[matches[1]] = immediate(v)
			}
			if v, err := strconv.Atoi(matches[3]); err == nil {
				circuit[matches[3]] = immediate(v)
			}
			switch matches[2] {
			case "AND":
				circuit[matches[4]] = and(matches[1], matches[3])
			case "OR":
				circuit[matches[4]] = or(matches[1], matches[3])
			case "LSHIFT":
				circuit[matches[4]] = lshift(matches[1], matches[3])
			case "RSHIFT":
				circuit[matches[4]] = rshift(matches[1], matches[3])
			}
		}
		if reNot.MatchString(line) {
			matches := reNot.FindStringSubmatch(line)
			if v, err := strconv.Atoi(matches[1]); err == nil {
				circuit[matches[1]] = immediate(v)
			}
			circuit[matches[2]] = not(matches[1])
		}
		if reVal.MatchString(line) {
			matches := reVal.FindStringSubmatch(line)
			if v, err := strconv.Atoi(matches[1]); err == nil {
				circuit[matches[1]] = immediate(v)
			}
			circuit[matches[2]] = val(matches[1])
		}
	}
	return circuit
}
